package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/suhailmuhammed157/simple_bank/api"
	"github.com/suhailmuhammed157/simple_bank/db_source"
)

const (
	driverName    = "postgres"
	dataSource    = "postgres://root:password@localhost:5433/simple_bank?sslmode=disable"
	serverAddress = "127.0.0.1:8080"
)

func main() {
	conn, err := sql.Open(driverName, dataSource)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	store := db_source.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start server")
	}

}
