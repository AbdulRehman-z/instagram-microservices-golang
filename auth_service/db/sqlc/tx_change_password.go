package db

import (
	"context"
	"database/sql"
	"fmt"
)

type ChangePasswordTxRequest struct {
	Email             string
	NewHashedPassword string
	AfterChange       func(user User) error
}

type ChangePasswordTxResult struct {
	User User
}

func (s *sqlStore) ChangePasswordTx(ctx context.Context, arg ChangePasswordTxRequest) (*ChangePasswordTxResult, error) {
	var result ChangePasswordTxResult

	err := s.execTx(ctx, sql.LevelDefault, func(q *Queries) error {

		var err error
		result.User, err = s.UpdateUser(ctx, UpdateUserParams{
			Email: arg.Email,
			HashedPassword: sql.NullString{
				Valid:  true,
				String: arg.NewHashedPassword,
			},
		})
		if err != nil {
			return fmt.Errorf("tx: failed to update user -- %w", err)
		}

		return arg.AfterChange(result.User)
	})

	if err != nil {
		return nil, err
	}

	return &result, nil
}
