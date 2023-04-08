package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
	APIGateway "github.com/muazwzxv/go-backend-masterclass/gateway/api"
	accountsHandler "github.com/muazwzxv/go-backend-masterclass/gateway/api/accounts"
	accountsModule "github.com/muazwzxv/go-backend-masterclass/modules/accounts"
	"github.com/muazwzxv/go-backend-masterclass/pkg"
	"go.uber.org/zap"
)

const (
	dbDriver      = "pgx"
	dbSource      = "postgresql://root:password@localhost:5432/go_masterclass?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	database, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	if err = database.Ping(); err != nil {
		log.Fatal(err)
	}
	store := db.NewStore(database)
	log, _ := zap.NewDevelopment()
	sugaredLogger := log.Sugar()

	/**
	  LOGGER should be in handler layer or module layer?
	*/

	server := pkg.NewServer(store, sugaredLogger)
	gateway := InitializeModules(server)
	gateway.Init(server.Mux)

	if err = server.Start(serverAddress); err != nil {
		sugaredLogger.Fatal("cannot start server: ", err)
	}
}

func InitializeModules(server *pkg.Server) *APIGateway.Gateway {
	accounts := accountsModule.New(server.Store, server.Log)
	accHandler := accountsHandler.New(accounts)

	return APIGateway.New(
		accHandler,
	)
}
