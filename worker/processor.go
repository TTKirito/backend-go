package worker

import (
	"context"

	db "github.com/TTKirito/backend-go/db/sqlc"
	"github.com/hibiken/asynq"
)

const (
	QueueCritical = "critical"
	QueueDefault  = "default"
)

type TaskProcessor interface {
	Start() error
	ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error
}

type RedistaskProcessor struct {
	server *asynq.Server
	store  db.Store
}

func NewRedistaskProcessor(redisOpt asynq.RedisClientOpt, store db.Store) TaskProcessor {
	server := asynq.NewServer(
		redisOpt,
		asynq.Config{
			Queues: map[string]int{
				QueueCritical: 10,
				QueueDefault:  5,
			},
		},
	)

	return &RedistaskProcessor{
		server: server,
		store:  store,
	}
}

func (processor *RedistaskProcessor) Start() error {
	mux := asynq.NewServeMux()
	mux.HandleFunc(TaskSendVerifyEmail, processor.ProcessTaskSendVerifyEmail)

	return processor.server.Start(mux)
}
