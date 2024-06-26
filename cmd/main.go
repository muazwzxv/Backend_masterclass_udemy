package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hibiken/asynq"
	_ "github.com/jackc/pgx/v5/stdlib"
	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
	APIGateway "github.com/muazwzxv/go-backend-masterclass/gateway/api"
	accountsHandler "github.com/muazwzxv/go-backend-masterclass/gateway/api/accounts"
	transfersHandler "github.com/muazwzxv/go-backend-masterclass/gateway/api/transfers"
	usersHandler "github.com/muazwzxv/go-backend-masterclass/gateway/api/users"
	"github.com/muazwzxv/go-backend-masterclass/gateway/rpc/user"
	accountsModule "github.com/muazwzxv/go-backend-masterclass/modules/accounts"
	transfersModule "github.com/muazwzxv/go-backend-masterclass/modules/transfers"
	adapter "github.com/muazwzxv/go-backend-masterclass/modules/transfers/adapters/accounts"
	usersModule "github.com/muazwzxv/go-backend-masterclass/modules/users"
	"github.com/muazwzxv/go-backend-masterclass/pb"
	"github.com/muazwzxv/go-backend-masterclass/pkg/authToken"
	"github.com/muazwzxv/go-backend-masterclass/pkg/config"
	"github.com/muazwzxv/go-backend-masterclass/pkg/rpcServer"
	"github.com/muazwzxv/go-backend-masterclass/pkg/server"
	"github.com/muazwzxv/go-backend-masterclass/pkg/worker"
	"github.com/muazwzxv/go-backend-masterclass/tools"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal("failed to load configs", err)
	}

	// connect to database
	database, err := sql.Open(cfg.DBDriver, cfg.DBSource)
	if err != nil {
		log.Fatal(err)
	}
	if err = database.Ping(); err != nil {
		log.Fatal(err)
	}
	store := db.NewStore(database)

	// connect to redis and setup worker
	redisOpt := asynq.RedisClientOpt{
		Addr:     cfg.RedisAddress,
		Password: cfg.RedisPassword,
	}
	taskDist := worker.NewRedisTaskDistributor(redisOpt)

	// setup logger
	log, _ := zap.NewDevelopment()
	sugaredLogger := log.Sugar()

	/**
		  LOGGER should be in handler layer or module layer?
	    RN IM PUTTING IT IN BOTH
	*/

	token, err := authToken.NewPaseto(cfg.TokenSymmetricKey)
	if err != nil {
		sugaredLogger.Fatal("failed to create token instance", err)
	}

	switch cfg.RunServer {
	case "RPC":
		runRpcServer(cfg, store, sugaredLogger, token, taskDist)
	case "HTTP":
		runHttpServer(cfg, store, sugaredLogger, token)
	default:
		go runTaskProcessor(redisOpt, store)
		go runGatewayServer(cfg, store, sugaredLogger, token)
		runRpcServer(cfg, store, sugaredLogger, token, taskDist)
	}
}

func runGatewayServer(cfg *config.Config, store *db.Store, log *zap.SugaredLogger, token authToken.IToken) {
	// Base RPC server dependency
	rpc := rpcServer.NewServer(rpcServer.ServerRequest{
		Cfg:   cfg,
		Store: store,
		Log:   log,
		Token: token,
	})

	// Setup services
	usersModule := usersModule.New(rpc.Config, rpc.Store, rpc.Log, rpc.Token)
	userService := user.NewUserServiceServer(rpc, usersModule)

	rpcMux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := pb.RegisterUserServiceHandlerServer(ctx, rpcMux, userService)
	if err != nil {
		log.Fatal("cannot register handler server")
	}

	mux := http.NewServeMux()
	mux.Handle("/", rpcMux)

	listener, err := net.Listen("tcp", cfg.HttpServerAddress)
	if err != nil {
		log.Fatal("cannot start rpc listener")
	}

	log.Info("start HTTP gateway server at %s", listener.Addr().String())

	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("failed to start HTTP gateway server")
	}
}

func runRpcServer(cfg *config.Config, store *db.Store, log *zap.SugaredLogger, token authToken.IToken, taskDistributor worker.TaskDistributor) {
	// Base RPC server dependency
	rpc := rpcServer.NewServer(rpcServer.ServerRequest{
		Cfg:             cfg,
		Store:           store,
		Log:             log,
		Token:           token,
		TaskDistributor: taskDistributor,
	})

	// Setup services
	usersModule := usersModule.New(rpc.Config, rpc.Store, rpc.Log, rpc.Token)
	userService := user.NewUserServiceServer(rpc, usersModule)

	// Setup gRPC server
	grpcServer := grpc.NewServer()

	// Register services to gRPC server
	pb.RegisterUserServiceServer(grpcServer, userService)

	listener, err := net.Listen("tcp", cfg.RpcServerAddress)
	if err != nil {
		log.Fatal("cannot start rpc listener")
	}

	log.Info("start gRPC server at %s", listener.Addr().String())

	reflection.Register(grpcServer)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("failed to start gRPC server")
	}
}

func runHttpServer(cfg *config.Config, store *db.Store, log *zap.SugaredLogger, token authToken.IToken) {
	server := server.NewServer(server.HttpServerRequest{
		Config: cfg,
		Store:  store,
		Log:    log,
		Token:  token,
	})
	gateway := InitializeModules(server)
	gateway.Init(server.Mux)

	if err := server.Start(cfg.HttpServerAddress); err != nil {
		// TODO: Implement graceful shutdown
		log.Fatal("cannot start server: ", err)
	}
}

func runTaskProcessor(redisOpt asynq.RedisClientOpt, store *db.Store) {
	processor := worker.NewRedisTaskProcessor(redisOpt, store)
	tools.Logger.Info("starting task processor ....")
	err := processor.Start()
	if err != nil {
		tools.Logger.Error("cannot start worker server", err)
	}
}

func InitializeModules(server *server.Server) *APIGateway.Gateway {
	accounts := accountsModule.New(server.Store, server.Log)
	accHandler := accountsHandler.New(accounts, server.Log)

	users := usersModule.New(server.Config, server.Store, server.Log, server.Token)
	usersHandler := usersHandler.New(users, server.Log)

	transfers := transfersModule.New(
		server.Store,
		server.Log,
		adapter.NewAccountsAdapter(accounts),
	)
	transfersHandler := transfersHandler.New(transfers, server.Log)

	return APIGateway.New(
		accHandler,
		usersHandler,
		transfersHandler,
	)
}
