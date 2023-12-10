package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
	VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (*VerifyEmailTxResult, error)
	RegisterUserTx(ctx context.Context, arg RegisterUserTXParams) (*RegisterUserTXResult, error)
	ChangePasswordTx(ctx context.Context, arg ChangePasswordTxRequest) (*ChangePasswordTxResult, error)
}

type sqlStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &sqlStore{
		Queries: New(db),
		db:      db,
	}
}

func (s *sqlStore) execTx(ctx context.Context, isolationLevel sql.IsolationLevel, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: isolationLevel,
	})
	if err != nil {
		return fmt.Errorf("tx: failed to begin transaction -- %w", err)
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("tx: rolling back -- %w", err)
		}
		return err
	}

	return tx.Commit()
}
