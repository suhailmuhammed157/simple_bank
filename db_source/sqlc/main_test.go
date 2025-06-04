package db_source

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/suhailmuhammed157/simple_bank/utils"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, configErr := utils.LoadConfig("../.")
	if configErr != nil {
		log.Fatal("Cannot load env")
	}
	var err error
	testDB, err = sql.Open(config.DBDriver, config.DataSource)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
