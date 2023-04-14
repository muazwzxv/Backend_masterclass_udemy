package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
	APIGateway "github.com/muazwzxv/go-backend-masterclass/gateway/api"
	accountsHandler "github.com/muazwzxv/go-backend-masterclass/gateway/api/accounts"
	accountsModule "github.com/muazwzxv/go-backend-masterclass/modules/accounts"
	usersModule "github.com/muazwzxv/go-backend-masterclass/modules/users"
	usersHandler "github.com/muazwzxv/go-backend-masterclass/gateway/api/users"
	"github.com/muazwzxv/go-backend-masterclass/pkg"
	"github.com/muazwzxv/go-backend-masterclass/pkg/config"
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
	server := pkg.NewServer(store, sugaredLogger)
	gateway := InitializeModules(server)
	gateway.Init(server.Mux)
  if err = server.Start(cfg.ServerAddress); err != nil {

		sugaredLogger.Fatal("cannot start server: ", err)
	}
}

func InitializeModules(server *pkg.Server) *APIGateway.Gateway {
	accounts := accountsModule.New(server.Store, server.Log)
	accHandler := accountsHandler.New(accounts, server.Log)

  users := usersModule.New(server.Store, server.Log)
  usersHandler := usersHandler.New(users, server.Log)

	return APIGateway.New(
		accHandler,
    usersHandler,
	)
}
