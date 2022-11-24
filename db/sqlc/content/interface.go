package db

import (
	"context"

	db "github.com/nina/db/sqlc/content/gen"
)

type Itf interface {
	CreateOneContent(ctx context.Context, arg db.CreateOneContentParams) (db.Content, error)
	GetAllContent(ctx context.Context) ([]db.Content, error)
	GetContentOneProject(ctx context.Context, projectName string) ([]db.Content, error)
	UpdateContent(ctx context.Context, arg db.UpdateContentParams) (db.Content, error)
	UpdateImageLink(ctx context.Context, arg db.UpdateImageLinkParams) (db.Content, error)
	UpdatePostStatus(ctx context.Context, arg db.UpdatePostStatusParams) error
}
