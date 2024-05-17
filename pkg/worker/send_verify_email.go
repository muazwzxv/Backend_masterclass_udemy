package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	"github.com/hibiken/asynq"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

const (
	TaskSendVerifyEmail = "task:send_verify_email"
)

type PayloadSendVerifyEmail struct {
	UserID int64 `json:"user_id"`
}

func (d *RedisTaskDistributor) SendVerifyEmail(ctx context.Context, payload *PayloadSendVerifyEmail, opts ...asynq.Option) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %v", err)
	}

	task := asynq.NewTask(TaskSendVerifyEmail, jsonPayload, opts...)
	info, err := d.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %v", err)
	}

	logger.Error("task type: %s \n payload: %+v \n queue: %s \n max retry: %d", task.Type(), task.Payload(), info.Queue, info.MaxRetry)

	return nil
}
