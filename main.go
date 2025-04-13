package main

import (
	"database/sql"
	"os"

	"net"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/suhailmuhammed157/simple_bank/db_source"
	"github.com/suhailmuhammed157/simple_bank/gapi"
	"github.com/suhailmuhammed157/simple_bank/pb"
	"github.com/suhailmuhammed157/simple_bank/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := utils.LoadConfig(".")

	if config.Env == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	if err != nil {

		log.Fatal().Msg("Config loading error")
	}

	conn, err := sql.Open(config.DBDriver, config.DataSource)
	if err != nil {
		log.Fatal().Msg("Cannot connect to database")
	}

	store := db_source.NewStore(conn)

	err = grpcServer(&config, store)
	if err != nil {
		log.Fatal().Msg("Cannot start server")
	}
}

func grpcServer(config *utils.Config, store *db_source.Store) error {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal().Msgf("Cannot start server %v", err)
	}

	//interceptor for logging
	grpcLogger := grpc.UnaryInterceptor(server.GrpcLogger)

	grpcServer := grpc.NewServer(grpcLogger)
	pb.RegisterSimpleBankServer(grpcServer, server)

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", config.GrpcServerAddress)
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	log.Info().Msgf("gRPC server started at %v", config.GrpcServerAddress)

	return grpcServer.Serve(lis)

}

// func apiServer(config *utils.Config, store *db_source.Store) error {
// 	server, err := api.NewServer(config, store)
// 	if err != nil {
// 		log.Fatal().Msgf("Cannot start server %v", err)
// 	}
// 	return server.Start(config.HttpServerAddress)
// }
