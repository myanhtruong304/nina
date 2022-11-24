package db

import (
	"context"

	db "github.com/nina/db/sqlc/twitter_bind_account/gen"
)

type Itf interface {
	CreateTwitterAccount(ctx context.Context, arg db.CreateTwitterAccountParams) (db.TwitterBindAccount, error)
}
