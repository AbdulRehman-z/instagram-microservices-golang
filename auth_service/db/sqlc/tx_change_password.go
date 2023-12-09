package db

import (
	"context"
	"database/sql"
	"time"
)

type ChangePasswordTxRequest struct {
	Email       string
	NewPassword string
}

type ChangePasswordTxResult struct {
	PasswordChangedAt time.Time
}

func (s *sqlStore) ChangePasswordTx(ctx context.Context, arg ChangePasswordTxRequest) (*ChangePasswordTxResult, error) {

	var result ChangePasswordTxResult

	err := s.execTx(ctx, sql.LevelDefault, func(q *Queries) error {

	})

}
