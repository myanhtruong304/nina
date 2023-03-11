-- name: AddUser :one
INSERT INTO users (
    username,
    hashed_pwd,
    email,
    pwd_changed_at,
    created_at
    ) VALUES ($1, $2, $3, $4, $5) RETURNING username, email;
