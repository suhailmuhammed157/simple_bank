package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/suhailmuhammed157/simple_bank/api"
	"github.com/suhailmuhammed157/simple_bank/db_source"
	"github.com/suhailmuhammed157/simple_bank/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Config loading error")
	}

	conn, err := sql.Open(config.DBDriver, config.DataSource)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	store := db_source.NewStore(conn)

	err = apiServer(config.HttpServerAddress, &config, store)
	if err != nil {
		log.Fatal("Cannot start server")
	}
}

func apiServer(httpAddress string, config *utils.Config, store *db_source.Store) error {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
	return server.Start(httpAddress)

}
