-- name: CreateProjectInfo :one
INSERT INTO projects_info (
    project_name,
    symbol,
    contract_address,
    explorer,
    twitter,
    facebook,
    linkedin,
    medium,
    telegram,
    website,
    git,
    cmc,
    coingecko) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING *;

-- name: GetOneProject :one
SELECT * FROM projects_info
WHERE project_name = $1
LIMIT 1;

-- name: GetListProjects :many
SELECT * FROM projects_info
ORDER BY project_name;

-- name: DeleteProject :exec
DELETE FROM projects_info WHERE project_name = $1;