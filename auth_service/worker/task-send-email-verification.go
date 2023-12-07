package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	db "github.com/AbdulRehman-z/instagram-microservices/auth_service/db/sqlc"
	"github.com/AbdulRehman-z/instagram-microservices/auth_service/mail"
	"github.com/AbdulRehman-z/instagram-microservices/auth_service/util"
	"github.com/hibiken/asynq"
)

type PayloadSendVerificationEmail struct {
	Email string
}

func (d *TaskDistributor) TaskSendVerificationEmail(ctx context.Context, payload *PayloadSendVerificationEmail, options ...asynq.Option) error {

	marshal, err := json.Marshal(&payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %v", err)
	}

	task := asynq.NewTask(
		TaskTypeEmailVerification,
		marshal,
		options...,
	)
	info, err := d.client.Enqueue(task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %v", err)
	}
	slog.Info("Enqueued task: ", "info", info)
	return nil
}

func (p *TaskProcessor) ProcessTask(ctx context.Context, task *asynq.Task) error {
	config, err := util.LoadConfig(".")
	if err != nil {
		return fmt.Errorf("cannot load config: %w", err)
	}

	var payload PayloadSendVerificationEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("cannot unmarshal payload: %w", err)
	}

	user, err := p.store.GetUser(ctx, payload.Email)
	if err != nil {
		return fmt.Errorf("cannot get user: %w", err)
	}

	verifyEmail, err := p.store.CreateVerifyEmail(ctx, db.CreateVerifyEmailParams{
		Username:   user.Username,
		Email:      user.Email,
		SecretCode: util.GenerateRandomString(6),
	})
	if err != nil {
		return fmt.Errorf("cannot create verify email: %w", err)
	}

	mail := mail.NewMailSender(config.EMAIL_USERNAME, config.EMAIL_FROM, config.APP_PASSWORD)
	receiverEmail := []string{"yousafbhaikhan10@gmail.com"}
	verificationUrl := fmt.Sprintf("Please verify your email by using the following link: http://localhost:8080/auth/verify_email?email_id=%d&secret_code=%s",
		verifyEmail.ID, verifyEmail.SecretCode)
	err = mail.SendEmail(receiverEmail, "Email Verification", verificationUrl)
	if err != nil {
		return fmt.Errorf("failed to send mail: %w", err)
	}

	return nil
}
