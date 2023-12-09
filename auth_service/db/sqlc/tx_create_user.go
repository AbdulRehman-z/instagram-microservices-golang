package db

import (
	"context"
	"database/sql"
	"fmt"
)

type RegisterUserTXParams struct {
	RegisterUserParams
	AfterRegister func(user User) error
}

type RegisterUserTXResult struct {
	User User
}

func (s *sqlStore) RegisterUserTx(ctx context.Context, arg RegisterUserTXParams) (*RegisterUserTXResult, error) {
	var result RegisterUserTXResult

	err := s.execTx(ctx, sql.LevelDefault, func(q *Queries) error {
		var err error
		result.User, err = s.RegisterUser(ctx, RegisterUserParams{
			Email:          arg.Email,
			HashedPassword: arg.HashedPassword,
		})
		if err != nil {
			return fmt.Errorf("tx: failed to register user -- %w", err)
		}

		return arg.AfterRegister(result.User)
	})

	if err != nil {
		return nil, err
	}

	return &result, nil
}
