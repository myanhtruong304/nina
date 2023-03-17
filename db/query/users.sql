-- name: AddUser :one
INSERT INTO users (
    username,
    hashed_pwd,
    user_email,
    pwd_changed_at,
    created_at
    ) VALUES ($1, $2, $3, $4, $5) RETURNING username, user_email;
