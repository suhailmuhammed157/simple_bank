package db_source

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/suhailmuhammed157/simple_bank/utils"
)

var testStore Store

func TestMain(m *testing.M) {
	config, configErr := utils.LoadConfig("../../.")
	if configErr != nil {
		log.Fatal("Cannot load env")
	}

	connPool, err := pgxpool.New(context.Background(), config.DataSource)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	testStore = NewStore(connPool)

	os.Exit(m.Run())
}
