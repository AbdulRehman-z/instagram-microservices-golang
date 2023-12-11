package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	db "github.com/AbdulRehman-z/instagram-microservices/auth_service/db/sqlc"
	"github.com/AbdulRehman-z/instagram-microservices/auth_service/util"
	"github.com/hibiken/asynq"
)

// TaskSignupVerificationEmail is the task type for sending verification email
type PayloadSendVerificationEmail struct {
	Email string
}

// TaskSendSignupEmail is the task type for sending signup email
func (d *TaskDistributor) TaskSendSignupEmail(ctx context.Context, payload *PayloadSendVerificationEmail, options ...asynq.Option) error {
	marshal, err := json.Marshal(&payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %v", err)
	}

	task := asynq.NewTask(
		TaskSignupVerificationEmail,
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

// TaskPasswordChangeVerificationEmail is the task type for sending password changed email
func (d *TaskDistributor) TaskPasswordChangeVerificationEmail(ctx context.Context, payload *PayloadSendVerificationEmail, options ...asynq.Option) error {
	marshal, err := json.Marshal(&payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %v", err)
	}

	task := asynq.NewTask(
		TaskPasswordChangeVerificationEmail,
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

// ProcessTask processes the task and sends the email
func (p *TaskProcessor) ProcessTask(ctx context.Context, task *asynq.Task) error {
	config, err := util.LoadConfig(".")
	if err != nil {
		return fmt.Errorf("cannot load config: %w", err)
	}

	fmt.Println("----------------------------")
	fmt.Println("task type: ", task.Type())
	fmt.Println("----------------------------")

	var payload PayloadSendVerificationEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("cannot unmarshal payload: %w", err)
	}

	user, err := p.store.GetUser(ctx, payload.Email)
	if err != nil {
		return fmt.Errorf("cannot get user: %w", err)
	}

	switch task.Type() {
	case TaskSignupVerificationEmail:
		err = sendSignupVerificationEmail(ctx, user, p, config)
		slog.Info("Sending signup verification email")
		if err != nil {
			return fmt.Errorf("err: %w", err)
		}
	case TaskPasswordChangeVerificationEmail:
		err = sendPasswordChangeVerificationEmail(ctx, user, p, config)
		slog.Info("Sending password change verification email")
		if err != nil {
			return fmt.Errorf("err: %w", err)
		}
	default:
		return fmt.Errorf("unexpected task type: %s", task.Type())
	}

	fmt.Println("----------------------------")
	fmt.Println("task processed successfully")
	fmt.Println("----------------------------")

	return nil
}

func sendSignupVerificationEmail(ctx context.Context, user db.User, p *TaskProcessor, config *util.Config) error {
	verifyEmail, err := p.store.CreateVerifyEmail(ctx, db.CreateVerifyEmailParams{
		Email:      user.Email,
		SecretCode: util.GenerateRandomString(6),
	})
	if err != nil {
		return fmt.Errorf("cannot create verify email: %w", err)
	}

	link := fmt.Sprintf("Please verify your email by using the following link:  http://localhost:8080/auth/verify-email?email_id=%d&secret_code=%s",
		verifyEmail.ID, verifyEmail.SecretCode)
	err = p.mailer.SendEmail([]string{user.Email}, "Verify your email", p.mailer.VerifyEmailTemplate(user.Email, link))
	if err != nil {
		return fmt.Errorf("cannot send email: %w", err)
	}

	return nil
}

func sendPasswordChangeVerificationEmail(ctx context.Context, user db.User, p *TaskProcessor, config *util.Config) error {
	verifyEmail, err := p.store.CreateVerifyEmail(ctx, db.CreateVerifyEmailParams{
		Email:      user.Email,
		SecretCode: util.GenerateRandomString(6),
	})
	if err != nil {
		return fmt.Errorf("cannot create verify email: %w", err)
	}

	link := fmt.Sprintf("Please verify your email by using the following link: http://localhost:8080/auth/verify-email?email_id=%d&secret_code=%s",
		verifyEmail.ID, verifyEmail.SecretCode)
	err = p.mailer.SendEmail([]string{user.Email}, "Verify your email", p.mailer.VerifyEmailTemplate(user.Email, link))
	if err != nil {
		return fmt.Errorf("cannot send email: %w", err)
	}

	return nil
}
