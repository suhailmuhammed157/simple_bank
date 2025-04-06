package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq"
	"github.com/suhailmuhammed157/simple_bank/api"
	"github.com/suhailmuhammed157/simple_bank/db_source"
	"github.com/suhailmuhammed157/simple_bank/gapi"
	"github.com/suhailmuhammed157/simple_bank/pb"
	"github.com/suhailmuhammed157/simple_bank/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	err = grpcServer(&config, store)
	if err != nil {
		log.Fatal("Cannot start server")
	}
}

func grpcServer(config *utils.Config, store *db_source.Store) error {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", config.GrpcServerAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("gRPC server started at %v", config.GrpcServerAddress)

	return grpcServer.Serve(lis)

}

func apiServer(config *utils.Config, store *db_source.Store) error {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
	return server.Start(config.HttpServerAddress)
}
