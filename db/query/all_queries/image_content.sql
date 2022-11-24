-- name: CreateOneImage :one
INSERT INTO image_content (
    project_name,
    image_content,
    content_id,
    created_at
) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateLink :one
UPDATE image_content SET link = $2, updated_at = $3
WHERE content_id = $1 RETURNING *;
