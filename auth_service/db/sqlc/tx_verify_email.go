package db

import (
	"context"
	"database/sql"
	"fmt"
)

type VerifyEmailTxParams struct {
	EmailId    int32
	SecretCode string
}

type VerifyEmailTxResult struct {
	User        User
	VerifyEmail VerifyEmail
}

func (s *sqlStore) VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (*VerifyEmailTxResult, error) {
	var result VerifyEmailTxResult

	err := s.execTx(ctx, sql.LevelDefault, func(q *Queries) error {
		var err error

		result.VerifyEmail, err = q.UpdateVerifyEmail(ctx, UpdateVerifyEmailParams{
			ID:         arg.EmailId,
			SecretCode: arg.SecretCode,
		})
		if err != nil {
			return fmt.Errorf("tx: failed to update verify_email -- %w", err)
		}

		result.User, err = q.UpdateUser(ctx, UpdateUserParams{
			Email: result.VerifyEmail.Email,
			IsEmailVerified: sql.NullBool{
				Valid: true,
				Bool:  true,
			},
		})
		if err != nil {
			return fmt.Errorf("tx: failed to update user -- %w", err)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &result, nil
}
