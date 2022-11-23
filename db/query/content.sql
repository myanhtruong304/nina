-- name: CreateOneContent :one
INSERT INTO content (
    project_name,
    content,
    char_count,
    image_link,
    platform,
    content_type,
    created_at,
    upload) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *;

-- name: GetContentOneProject :many
SELECT * FROM content
WHERE project_name = $1
ORDER BY created_at;

-- name: GetAllContent :many
SELECT * FROM content
ORDER BY project_name;

-- name: UpdateImageLink :one
UPDATE content SET image_link = $2, updated_at = $3
WHERE id = $1 RETURNING *;

-- name: UpdateContent :one
UPDATE content SET content = $2
WHERE id = $1
RETURNING *;

-- name: UpdatePostStatus :exec
UPDATE content SET upload = $2, updated_at = $3
WHERE id = $1;