package worker

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	db "github.com/muazwzxv/go-backend-masterclass/db/sqlc"
	"github.com/muazwzxv/go-backend-masterclass/tools"
)

type TaskProcessor interface {
	SendVerifyEmail(ctx context.Context, task *asynq.Task) error
	Start() error
}

type RedisTaskProcessor struct {
	server *asynq.Server
	store  *db.Store
}

func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt, db *db.Store) TaskProcessor {
	svr := asynq.NewServer(redisOpt, asynq.Config{
		Concurrency: 10,
	})

	return &RedisTaskProcessor{
		server: svr,
		store:  db,
	}
}

func (t *RedisTaskProcessor) SendVerifyEmail(ctx context.Context, task *asynq.Task) error {
	var payload PayloadSendVerifyEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	user, err := t.store.GetUser(ctx, payload.UserID)
	if err != nil {
		if err != sql.ErrNoRows {
			return fmt.Errorf("user not found: %w", err)
		}
		return fmt.Errorf("failed to get user data: %w", err)
	}

	// TODO: send email
	tools.Logger.Info("Sending email verification for user ID: %d \n First Name: %s", user.ID, user.FirstName)

	return nil
}

func (t *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()

	mux.HandleFunc(TaskSendVerifyEmail, t.SendVerifyEmail)

	return t.server.Start(mux)
}
