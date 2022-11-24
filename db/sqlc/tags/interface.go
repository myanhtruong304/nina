package db

import (
	"context"

	db "github.com/nina/db/sqlc/tags/gen"
)

type Itf interface {
	CreateOneTag(ctx context.Context, arg db.CreateOneTagParams) (db.Tag, error)
	UpdateProjectTag(ctx context.Context, arg db.UpdateProjectTagParams) (db.Tag, error)
}
