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
	dataSource = "postgres://root:password@localhost:5433/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(driverName, dataSource)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
