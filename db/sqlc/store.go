package db

import "database/sql"

type Store struct {
	Querier
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		Querier: &SQLStore{
			Queries: New(db),
			db:      db,
		},
	}
}