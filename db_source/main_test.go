package db_source

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	driverName = "postgres"
	dataSource = "postgres://root:password@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(driverName, dataSource)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
