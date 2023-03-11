// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: users.sql

package db

import (
	"context"
)

const addUser = `-- name: AddUser :one
INSERT INTO users (
    username,
    hashed_pwd,
    email,
    pwd_changed_at,
    created_at
    ) VALUES ($1, $2, $3, $4, $5) RETURNING username, email
`

type AddUserParams struct {
	Username     string      `json:"username"`
	HashedPwd    string      `json:"hashed_pwd"`
	Email        string      `json:"email"`
	PwdChangedAt interface{} `json:"pwd_changed_at"`
	CreatedAt    interface{} `json:"created_at"`
}

type AddUserRow struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (q *Queries) AddUser(ctx context.Context, arg AddUserParams) (AddUserRow, error) {
	row := q.queryRow(ctx, q.addUserStmt, addUser,
		arg.Username,
		arg.HashedPwd,
		arg.Email,
		arg.PwdChangedAt,
		arg.CreatedAt,
	)
	var i AddUserRow
	err := row.Scan(&i.Username, &i.Email)
	return i, err
}