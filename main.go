package main

import (
	"context"
	"os"

	"net"

	"github.com/hibiken/asynq"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	db_source "github.com/suhailmuhammed157/simple_bank/db_source/sqlc"
	"github.com/suhailmuhammed157/simple_bank/email"
	"github.com/suhailmuhammed157/simple_bank/gapi"
	"github.com/suhailmuhammed157/simple_bank/pb"
	"github.com/suhailmuhammed157/simple_bank/utils"
	"github.com/suhailmuhammed157/simple_bank/worker"
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

	// conn, err := sql.Open(config.DBDriver, config.DataSource)
	conn, err := pgxpool.New(context.Background(), config.DataSource)
	if err != nil {
		log.Fatal().Msg("Cannot connect to database")
	}

	store := db_source.NewStore(conn)

	redisConn := asynq.RedisClientOpt{Addr: config.RedisServerAddress}
	taskDistributor := worker.NewRedisTaskDistributor(redisConn)

	go runTaskProcessor(store, redisConn, &config)
	err = grpcServer(&config, store, taskDistributor)
	if err != nil {
		log.Fatal().Msg("Cannot start server")
	}
}

func runTaskProcessor(store db_source.Store, redisConn asynq.RedisClientOpt, config *utils.Config) {
	mailer := email.NewEmailSender(config.EmailHost, config.EmailPort, config.EmailUser, config.EmailPassword)
	taskProcessor := worker.NewRedisTaskProcessor(redisConn, store, mailer)
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Msg("Cannot start server")
	}
}

func grpcServer(config *utils.Config, store db_source.Store, taskDistributor worker.TaskDistributor) error {
	server, err := gapi.NewServer(config, store, taskDistributor)
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
