// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package postgres

import (
	"context"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users (name)
VALUES ($1)
`

func (q *Queries) CreateUser(ctx context.Context, name string) error {
	_, err := q.db.ExecContext(ctx, createUser, name)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
UPDATE users
SET deleted_at = NOW()
WHERE id = $1
  AND deleted_at IS NULL
`

func (q *Queries) DeleteUser(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, name, created_at, updated_at, deleted_at
FROM users
WHERE id = $1
  AND deleted_at IS NULL
`

func (q *Queries) GetUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, name, created_at, updated_at, deleted_at
FROM users
WHERE deleted_at IS NULL
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET name = $2
WHERE id = $1
  AND deleted_at IS NULL
`

type UpdateUserParams struct {
	ID   string
	Name string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser, arg.ID, arg.Name)
	return err
}
