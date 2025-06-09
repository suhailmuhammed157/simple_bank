package worker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	db_source "github.com/suhailmuhammed157/simple_bank/db_source/sqlc"
	"github.com/suhailmuhammed157/simple_bank/utils"
)

const TaskSendVerifyEmail = "task:send_verify_email"

type PayloadSendVerifyEmail struct {
	Username string `json:"username"`
}

func (taskDistributor RedisTaskDistributor) DistributeTaskSendVerifyEmail(ctx context.Context, payload *PayloadSendVerifyEmail, opts ...asynq.Option) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshall task payload: %w", err)
	}
	task := asynq.NewTask(TaskSendVerifyEmail, jsonPayload, opts...)

	info, err := taskDistributor.client.Enqueue(task, opts...)

	if err != nil {
		return fmt.Errorf("failed to enqueue task: %w", err)
	}
	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).Str("Queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("enqueued task")
	return nil
}

func (taskProcessor RedisTaskProcessor) ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error {
	var p PayloadSendVerifyEmail
	if err := json.Unmarshal(task.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	user, err := taskProcessor.store.GetUser(ctx, p.Username)
	if err != nil {
		if err == db_source.NoRowFound {
			return fmt.Errorf("user does not exists: %v: %w", err, asynq.SkipRetry) // no need to retry
		}

		return fmt.Errorf("failed to get user: %w", err) // need to retry
	}

	verifyEmail, err := taskProcessor.store.CreateVerifyEmail(ctx, db_source.CreateVerifyEmailParams{
		Username:   user.Username,
		Email:      user.Email,
		SecretCode: utils.RandomString(4),
	})
	if err != nil {
		return fmt.Errorf("failed to create verify email: %w", err)
	}
	taskProcessor.mailer.SendEmail("register@simplebank.com", user.Email, "Welcome to Simple Bank", verifyEmail.SecretCode)
	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).Str("email", user.Email).Msg("processed task")

	return nil
}
