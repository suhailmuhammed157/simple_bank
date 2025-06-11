package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"

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
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	WebPort          = 8080
	InternalWebPort  = 8081
	interruptSignals = []os.Signal{
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGINT,
	}
)

func main() {
	config, err := utils.LoadConfig(".")

	if config.Env == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	if err != nil {
		log.Fatal().Msg("Config loading error")
	}

	// listen to these many interrupt signals
	notifyGracefulShutdownSignalCtx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	conn, err := pgxpool.New(notifyGracefulShutdownSignalCtx, config.DataSource)
	if err != nil {
		log.Fatal().Msg("Cannot connect to database")
	}

	store := db_source.NewStore(conn)

	redisConn := asynq.RedisClientOpt{Addr: config.RedisServerAddress}

	taskDistributor := worker.NewRedisTaskDistributor(redisConn)

	waitGroup, notifyGracefulShutdownSignalCtx := errgroup.WithContext(notifyGracefulShutdownSignalCtx)

	runTaskProcessor(waitGroup, notifyGracefulShutdownSignalCtx, store, redisConn, &config)
	grpcServer(waitGroup, notifyGracefulShutdownSignalCtx, &config, store, taskDistributor)

	err = waitGroup.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg("error from wait group ")
	}

}

func runTaskProcessor(waitGroup *errgroup.Group, notifyGracefulShutdownSignalCtx context.Context, store db_source.Store, redisConn asynq.RedisClientOpt, config *utils.Config) {
	mailer := email.NewEmailSender(config.EmailHost, config.EmailPort, config.EmailUser, config.EmailPassword)
	taskProcessor := worker.NewRedisTaskProcessor(redisConn, store, mailer)
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Msg("Cannot start server")
	}

	waitGroup.Go(func() error {
		<-notifyGracefulShutdownSignalCtx.Done()
		log.Info().Msgf("graceful shutdown task processor ")
		taskProcessor.Shutdown()
		log.Info().Msgf("task processor stopped")
		return nil
	})
}

func grpcServer(waitGroup *errgroup.Group, notifyGracefulShutdownSignalCtx context.Context, config *utils.Config, store db_source.Store, taskDistributor worker.TaskDistributor) {
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

	//go routine for the server
	waitGroup.Go(func() error {
		log.Info().Msgf("gRPC server started at %v", config.GrpcServerAddress)
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Fatal().Msg("Cannot start server")
			return err
		}
		return nil
	})

	//go routine to listen for ctx when os signals to interrupt
	waitGroup.Go(func() error {
		<-notifyGracefulShutdownSignalCtx.Done()
		log.Info().Msgf("graceful shutdown gRPC server ")
		grpcServer.GracefulStop()
		log.Info().Msgf("gRPC server stopped")
		return nil
	})

}

// func apiServer(config *utils.Config, store *db_source.Store) error {
// 	server, err := api.NewServer(config, store)
// 	if err != nil {
// 		log.Fatal().Msgf("Cannot start server %v", err)
// 	}
// 	return server.Start(config.HttpServerAddress)
// }
