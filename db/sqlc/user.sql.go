// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users (
    username,
    hashed_password,
    full_name,
    email
) VALUES (
    ?, ?, ?, ?
    )
`

type CreateUserParams struct {
	Username       string
	HashedPassword string
	FullName       string
	Email          string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser,
		arg.Username,
		arg.HashedPassword,
		arg.FullName,
		arg.Email,
	)
	return err
}

const getUser = `-- name: GetUser :one
SELECT username, hashed_password, full_name, email, is_email_verified, password_changed_at, created_at FROM users
WHERE username = ? LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.IsEmailVerified,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET
    hashed_password = COALESCE(?, hashed_password),
    password_changed_at = COALESCE(?, password_changed_at),
    full_name = COALESCE(?, full_name),
    email = COALESCE(?, email),
    is_email_verified = COALESCE(?, is_email_verified)
WHERE
    username = ?
`

type UpdateUserParams struct {
	HashedPassword    sql.NullString
	PasswordChangedAt sql.NullTime
	FullName          sql.NullString
	Email             sql.NullString
	IsEmailVerified   sql.NullBool
	Username          string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.HashedPassword,
		arg.PasswordChangedAt,
		arg.FullName,
		arg.Email,
		arg.IsEmailVerified,
		arg.Username,
	)
	return err
}
