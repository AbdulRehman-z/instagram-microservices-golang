package worker

import (
	"context"
	"fmt"
	"log/slog"

	db "github.com/AbdulRehman-z/instagram-microservices/auth_service/db/sqlc"
	"github.com/AbdulRehman-z/instagram-microservices/auth_service/mail"
	"github.com/hibiken/asynq"
)

var (
	CriticaLQueue                       = "critical"
	DefaultQueue                        = "default"
	TaskSignupVerificationEmail         = "signup_verification"
	TaskPasswordChangeVerificationEmail = "password_changed"
)

type Processor interface {
	Start() error
	ProcessTask(ctx context.Context, task *asynq.Task) error
}

type TaskProcessor struct {
	server *asynq.Server
	store  db.Store
	mailer mail.Mailer
}

func NewProcessor(options asynq.RedisClientOpt, store db.Store, mailer mail.Mailer) Processor {
	server := asynq.NewServer(options, asynq.Config{
		Queues: map[string]int{
			CriticaLQueue: 6,
			DefaultQueue:  3,
		},
		ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
			slog.Error("Error processing task from queue: ", "err", err, "type", task.Type, "payload", task.Payload)
		}),
	})

	return &TaskProcessor{
		server: server,
		store:  store,
		mailer: mailer,
	}
}

func (p *TaskProcessor) Start() error {
	fmt.Println("----------------------------------")
	fmt.Println("Starting task processor...")
	fmt.Println("----------------------------------")

	mux := asynq.NewServeMux()
	mux.HandleFunc(TaskSignupVerificationEmail, p.ProcessTask)
	mux.HandleFunc(TaskPasswordChangeVerificationEmail, p.ProcessTask)
	return p.server.Start(mux)
}
