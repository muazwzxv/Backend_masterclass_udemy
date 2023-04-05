package db_test

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"testing"

	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
)

const (
	dbDriver = "pgx"
	dbSource = "postgresql://root:password@localhost:5432/go_masterclass?sslmode=disable"
)

var testQueries *db.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal(err)
	}

	testDB = conn
	testQueries = db.New(testDB)
	m.Run()
}
