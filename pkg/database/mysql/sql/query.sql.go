// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package db

import (
	"context"
	"database/sql"
)

const getUser = `-- name: GetUser :one
SELECT id, name, property
FROM users
WHERE id = ? and property = ?
`

type GetUserParams struct {
	ID       string `json:"id"`
	Property string `json:"property"`
}

func (q *Queries) GetUser(ctx context.Context, arg GetUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, arg.ID, arg.Property)
	var i User
	err := row.Scan(&i.ID, &i.Name, &i.Property)
	return i, err
}

const persistUser = `-- name: PersistUser :execresult
INSERT INTO users (id, name, property) VALUES (?, ?, ?)
`

type PersistUserParams struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Property string `json:"property"`
}

func (q *Queries) PersistUser(ctx context.Context, arg PersistUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, persistUser, arg.ID, arg.Name, arg.Property)
}
