package db

import "database/sql"

type Store interface {
	Querier
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
