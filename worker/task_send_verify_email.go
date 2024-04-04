package worker

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
)

const TaskSendVerifyEmail = "task:send_verify_email"

type PayloadSendVerifyEmail struct {
	Username string `json:"username"`
}

func (distributor *RedisTaskDistributor) DistributeTaskSendVerifyEmail(
	ctx context.Context,
	payload *PayloadSendVerifyEmail,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %w", err)
	}

	task := asynq.NewTask(TaskSendVerifyEmail, jsonPayload, opts...)
	_, err = distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %w", err)
	}

	log.Printf("sendQUEUE:")
	return nil
}

func (processor *RedistaskProcessor) ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error {
	var payload PayloadSendVerifyEmail

	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload %w", asynq.SkipRetry)
	}

	user, err := processor.store.GetUser(ctx, payload.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user doesn't exist: %w", asynq.SkipRetry)
		}

		return fmt.Errorf("failed to get user: %w", err)
	}
	log.Println(user)

	return nil
}
