package db

import (
	"context"

	db "github.com/nina/db/sqlc/image_content/gen"
)

type Itf interface {
	CreateOneImage(ctx context.Context, arg db.CreateOneImageParams) (db.ImageContent, error)
	UpdateLink(ctx context.Context, arg db.UpdateLinkParams) (db.ImageContent, error)
}
