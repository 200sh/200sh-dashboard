// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: user.sql

package repository

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO user (provider_id, provider, name, email, status)
VALUES (?, ?, ?, ?, ?)
RETURNING id, provider_id, provider, name, email, status, created_at, updated_at
`

type CreateUserParams struct {
	ProviderID string `json:"provider_id"`
	Provider   string `json:"provider"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Status     int64  `json:"status"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ProviderID,
		arg.Provider,
		arg.Name,
		arg.Email,
		arg.Status,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.ProviderID,
		&i.Provider,
		&i.Name,
		&i.Email,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findUserByProviderID = `-- name: FindUserByProviderID :one
SELECT id, provider_id, provider, name, email, status, created_at, updated_at
FROM user
WHERE provider_id = ?
LIMIT 1
`

func (q *Queries) FindUserByProviderID(ctx context.Context, providerID string) (User, error) {
	row := q.db.QueryRowContext(ctx, findUserByProviderID, providerID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.ProviderID,
		&i.Provider,
		&i.Name,
		&i.Email,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE user
SET name = ?,
    email = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?
`

type UpdateUserParams struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	ID    int64  `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser, arg.Name, arg.Email, arg.ID)
	return err
}
