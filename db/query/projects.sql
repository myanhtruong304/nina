-- name: AddProject :one
INSERT INTO projects (
    project_name,
    symbol,
    website,
    twitter,
    telegram,
    facebook,
    linkedin,
    medium,
    whitepaper,
    email,
    contract_address,
    explorer,
    owner,
    created_at
    ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING project_name;

-- name: GetProject :one
SELECT * FROM projects
WHERE project_name = $1
LIMIT 1;

-- name: GetAllProject :many
SELECT * FROM projects
ORDER BY project_name;

-- name: ModifyProjectInfo :one
UPDATE projects SET
    symbol = $2,
    website = $3,
    twitter = $4,
    telegram = $5,
    facebook = $6,
    linkedin = $7,
    medium = $8,
    whitepaper = $9,
    email = $10,
    contract_address = $11,
    explorer = $12
WHERE project_name = $1 RETURNING *;