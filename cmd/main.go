package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
	APIGateway "github.com/muazwzxv/go-backend-masterclass/gateway/api"
	accountsHandler "github.com/muazwzxv/go-backend-masterclass/gateway/api/accounts"
	transfersHandler "github.com/muazwzxv/go-backend-masterclass/gateway/api/transfers"
	usersHandler "github.com/muazwzxv/go-backend-masterclass/gateway/api/users"
	accountsModule "github.com/muazwzxv/go-backend-masterclass/modules/accounts"
	transfersModule "github.com/muazwzxv/go-backend-masterclass/modules/transfers"
	adapter "github.com/muazwzxv/go-backend-masterclass/modules/transfers/adapters/accounts"
	usersModule "github.com/muazwzxv/go-backend-masterclass/modules/users"
	"github.com/muazwzxv/go-backend-masterclass/pkg/authToken"
	"github.com/muazwzxv/go-backend-masterclass/pkg/config"
	"github.com/muazwzxv/go-backend-masterclass/pkg/server"
	"go.uber.org/zap"
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

	// setup logger
	log, _ := zap.NewDevelopment()
	sugaredLogger := log.Sugar()

	/**
		  LOGGER should be in handler layer or module layer?
	    RN IM PUTTING IT IN BOTH
	*/

	// TODO - Put symmetric key in config file
	token, err := authToken.NewPaseto(cfg.TokenSymmetricKey)
	if err != nil {
		sugaredLogger.Fatal("failed to create token instance", err)
	}

	server := server.NewServer(cfg, store, sugaredLogger, token)
	gateway := InitializeModules(server)
	gateway.Init(server.Mux)

	if err = server.Start(cfg.ServerAddress); err != nil {
		// TODO: Implement graceful shutdown
		sugaredLogger.Fatal("cannot start server: ", err)
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
