package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/muazwzxv/go-backend-masterclass/api"
	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
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

	server := api.NewServer(store, sugaredLogger)
	server.SetupRoutes()

	if err = server.Start(serverAddress); err != nil {
    sugaredLogger.Fatal("cannot start server: ", err)
  }
}
