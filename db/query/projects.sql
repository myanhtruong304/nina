-- name: AddProject :one
INSERT INTO project (
    project_name,
    symbol,
    contract_address,
    owner,
    created_at
    ) VALUES ($1, $2, $3, $4, $5) RETURNING project_name;

-- name: GetProject :one
SELECT * FROM project
WHERE project_name = $1
LIMIT 1;

-- name: GetAllProject :many
SELECT * FROM project
ORDER BY project_name;