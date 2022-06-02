// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package db

import (
	"context"
	"database/sql"
)

const getUser = `-- name: GetUser :one
SELECT id, name
FROM users
WHERE id = ? and name = ?
`

type GetUserParams struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) GetUser(ctx context.Context, arg GetUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, arg.ID, arg.Name)
	var i User
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const persistUser = `-- name: PersistUser :execresult
INSERT INTO users (id, name) VALUES (?, ?)
`

type PersistUserParams struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) PersistUser(ctx context.Context, arg PersistUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, persistUser, arg.ID, arg.Name)
}