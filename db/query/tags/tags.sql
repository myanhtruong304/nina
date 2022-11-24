-- name: CreateOneTag :one
INSERT INTO tags (
    tag,
    project_name
) VALUES ($1, $2) RETURNING *;

-- name: UpdateProjectTag :one
UPDATE tags SET project_name = $2 WHERE tag = $1 RETURNING *;