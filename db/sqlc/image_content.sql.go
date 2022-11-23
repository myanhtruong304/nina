// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: image_content.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createOneImage = `-- name: CreateOneImage :one
INSERT INTO image_content (
    project_name,
    image_content,
    content_id,
    created_at
) VALUES ($1, $2, $3, $4) RETURNING id, project_name, image_content, content_id, created_at, updated_at, link
`

type CreateOneImageParams struct {
	ProjectName  string    `json:"project_name"`
	ImageContent string    `json:"image_content"`
	ContentID    int32     `json:"content_id"`
	CreatedAt    time.Time `json:"created_at"`
}

func (q *Queries) CreateOneImage(ctx context.Context, arg CreateOneImageParams) (ImageContent, error) {
	row := q.queryRow(ctx, q.createOneImageStmt, createOneImage,
		arg.ProjectName,
		arg.ImageContent,
		arg.ContentID,
		arg.CreatedAt,
	)
	var i ImageContent
	err := row.Scan(
		&i.ID,
		&i.ProjectName,
		&i.ImageContent,
		&i.ContentID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Link,
	)
	return i, err
}

const updateLink = `-- name: UpdateLink :one
UPDATE image_content SET link = $2, updated_at = $3
WHERE content_id = $1 RETURNING id, project_name, image_content, content_id, created_at, updated_at, link
`

type UpdateLinkParams struct {
	ContentID int32          `json:"content_id"`
	Link      sql.NullString `json:"link"`
	UpdatedAt time.Time      `json:"updated_at"`
}

func (q *Queries) UpdateLink(ctx context.Context, arg UpdateLinkParams) (ImageContent, error) {
	row := q.queryRow(ctx, q.updateLinkStmt, updateLink, arg.ContentID, arg.Link, arg.UpdatedAt)
	var i ImageContent
	err := row.Scan(
		&i.ID,
		&i.ProjectName,
		&i.ImageContent,
		&i.ContentID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Link,
	)
	return i, err
}
