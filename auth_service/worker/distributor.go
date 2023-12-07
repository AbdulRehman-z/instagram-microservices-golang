package worker

import (
	"context"

	"github.com/hibiken/asynq"
)

type Distributor interface {
	TaskSendVerificationEmail(ctx context.Context, payload *PayloadSendVerificationEmail, options ...asynq.Option) error
}

type TaskDistributor struct {
	client *asynq.Client
}

func NewDistributor(options *asynq.RedisClientOpt) Distributor {
	client := asynq.NewClient(options)
	return &TaskDistributor{
		client: client,
	}
}
