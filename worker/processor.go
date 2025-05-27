package worker

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/suhailmuhammed157/simple_bank/db_source"
)

type TaskProcessor interface {
	ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error
	Start() error
}

type RedisTaskProcessor struct {
	server *asynq.Server
	store  *db_source.Store
}

func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt, store *db_source.Store) TaskProcessor {
	server := asynq.NewServer(
		redisOpt,
		asynq.Config{},
	)

	return &RedisTaskProcessor{
		server: server,
		store:  store,
	}
}

func (redisTaskProcessor *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()
	mux.HandleFunc(TaskSendVerifyEmail, redisTaskProcessor.ProcessTaskSendVerifyEmail)
	return redisTaskProcessor.server.Start(mux)
}
