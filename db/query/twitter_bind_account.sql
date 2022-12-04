-- name: CreateTwitterAccount :one
INSERT INTO twitter_bind_account (
    project_name,
    access_token,
    access_token_secret
) VALUES ($1, $2, $3) RETURNING *;