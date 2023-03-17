-- name: AddContent :one
INSERT INTO contents (
    id,
    content,
    word_count,
    content_type,
    project_name,
    image_text,
    created_at
    ) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;

-- name: UpdateSchedule :one
UPDATE contents SET
    schedule_time = $2
WHERE id = $1 RETURNING schedule_time;

-- name: UpdateProgress :one
UPDATE contents SET
    facebook_check = $2,
    twitter_check = $3,
    linkedin_check = $4
WHERE id = $1 RETURNING id;

-- name: UpdateImageLink :one
UPDATE contents SET
    image_link = $2,
    last_updated_at = $3
WHERE id = $1 RETURNING image_link;