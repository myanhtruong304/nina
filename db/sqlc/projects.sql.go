// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: projects.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const addProject = `-- name: AddProject :one
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
    ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING project_name
`

type AddProjectParams struct {
	ProjectName     string         `json:"project_name"`
	Symbol          sql.NullString `json:"symbol"`
	Website         sql.NullString `json:"website"`
	Twitter         sql.NullString `json:"twitter"`
	Telegram        sql.NullString `json:"telegram"`
	Facebook        sql.NullString `json:"facebook"`
	Linkedin        sql.NullString `json:"linkedin"`
	Medium          sql.NullString `json:"medium"`
	Whitepaper      sql.NullString `json:"whitepaper"`
	Email           sql.NullString `json:"email"`
	ContractAddress sql.NullString `json:"contract_address"`
	Explorer        sql.NullString `json:"explorer"`
	Owner           string         `json:"owner"`
	CreatedAt       time.Time      `json:"created_at"`
}

func (q *Queries) AddProject(ctx context.Context, arg AddProjectParams) (string, error) {
	row := q.queryRow(ctx, q.addProjectStmt, addProject,
		arg.ProjectName,
		arg.Symbol,
		arg.Website,
		arg.Twitter,
		arg.Telegram,
		arg.Facebook,
		arg.Linkedin,
		arg.Medium,
		arg.Whitepaper,
		arg.Email,
		arg.ContractAddress,
		arg.Explorer,
		arg.Owner,
		arg.CreatedAt,
	)
	var project_name string
	err := row.Scan(&project_name)
	return project_name, err
}

const getAllProject = `-- name: GetAllProject :many
SELECT project_name, symbol, contract_address, explorer, website, twitter, facebook, linkedin, telegram, medium, whitepaper, email, owner, created_at, last_updated_at FROM projects
ORDER BY project_name
`

func (q *Queries) GetAllProject(ctx context.Context) ([]Project, error) {
	rows, err := q.query(ctx, q.getAllProjectStmt, getAllProject)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Project
	for rows.Next() {
		var i Project
		if err := rows.Scan(
			&i.ProjectName,
			&i.Symbol,
			&i.ContractAddress,
			&i.Explorer,
			&i.Website,
			&i.Twitter,
			&i.Facebook,
			&i.Linkedin,
			&i.Telegram,
			&i.Medium,
			&i.Whitepaper,
			&i.Email,
			&i.Owner,
			&i.CreatedAt,
			&i.LastUpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProject = `-- name: GetProject :one
SELECT project_name, symbol, contract_address, explorer, website, twitter, facebook, linkedin, telegram, medium, whitepaper, email, owner, created_at, last_updated_at FROM projects
WHERE project_name = $1
LIMIT 1
`

func (q *Queries) GetProject(ctx context.Context, projectName string) (Project, error) {
	row := q.queryRow(ctx, q.getProjectStmt, getProject, projectName)
	var i Project
	err := row.Scan(
		&i.ProjectName,
		&i.Symbol,
		&i.ContractAddress,
		&i.Explorer,
		&i.Website,
		&i.Twitter,
		&i.Facebook,
		&i.Linkedin,
		&i.Telegram,
		&i.Medium,
		&i.Whitepaper,
		&i.Email,
		&i.Owner,
		&i.CreatedAt,
		&i.LastUpdatedAt,
	)
	return i, err
}

const modifyProjectInfo = `-- name: ModifyProjectInfo :one
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
WHERE project_name = $1 RETURNING project_name, symbol, contract_address, explorer, website, twitter, facebook, linkedin, telegram, medium, whitepaper, email, owner, created_at, last_updated_at
`

type ModifyProjectInfoParams struct {
	ProjectName     string         `json:"project_name"`
	Symbol          sql.NullString `json:"symbol"`
	Website         sql.NullString `json:"website"`
	Twitter         sql.NullString `json:"twitter"`
	Telegram        sql.NullString `json:"telegram"`
	Facebook        sql.NullString `json:"facebook"`
	Linkedin        sql.NullString `json:"linkedin"`
	Medium          sql.NullString `json:"medium"`
	Whitepaper      sql.NullString `json:"whitepaper"`
	Email           sql.NullString `json:"email"`
	ContractAddress sql.NullString `json:"contract_address"`
	Explorer        sql.NullString `json:"explorer"`
}

func (q *Queries) ModifyProjectInfo(ctx context.Context, arg ModifyProjectInfoParams) (Project, error) {
	row := q.queryRow(ctx, q.modifyProjectInfoStmt, modifyProjectInfo,
		arg.ProjectName,
		arg.Symbol,
		arg.Website,
		arg.Twitter,
		arg.Telegram,
		arg.Facebook,
		arg.Linkedin,
		arg.Medium,
		arg.Whitepaper,
		arg.Email,
		arg.ContractAddress,
		arg.Explorer,
	)
	var i Project
	err := row.Scan(
		&i.ProjectName,
		&i.Symbol,
		&i.ContractAddress,
		&i.Explorer,
		&i.Website,
		&i.Twitter,
		&i.Facebook,
		&i.Linkedin,
		&i.Telegram,
		&i.Medium,
		&i.Whitepaper,
		&i.Email,
		&i.Owner,
		&i.CreatedAt,
		&i.LastUpdatedAt,
	)
	return i, err
}
