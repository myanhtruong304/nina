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

-- name: UpdateProjectInfo :one
UPDATE projects_info SET symbol = $2, contract_address = $3, explorer = $4, twitter = $5, facebook = $6, linkedin = $7, medium = $8, telegram = $9, website = $10, git = $11, cmc = $12, coingecko = $13
WHERE project_name = $1 RETURNING *;