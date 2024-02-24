// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (email, 
password) 
VALUES ($1, $2) 
RETURNING id, email, create_at, update_at
`

type CreateUserParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRow struct {
	ID       uuid.UUID    `json:"id"`
	Email    string       `json:"email"`
	CreateAt time.Time    `json:"create_at"`
	UpdateAt sql.NullTime `json:"update_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Email, arg.Password)
	var i CreateUserRow
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, password, create_at, update_at FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}