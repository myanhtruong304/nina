package entities

import (
	"database/sql"
	"log"

	db "github.com/nina/db/sqlc"
	"github.com/nina/util"
)

type Entity struct {
	repo db.Querier
	cfg  util.Config
}

func NewEntity(cfg util.Config) *Entity {
	conn, err := sql.Open(cfg.DBDriver, cfg.DBSource)
	if err != nil {
		log.Fatal("[NewEntity(cfg *util.Config)] can not load config", err)
	}

	store := db.NewStore(conn)
	return &Entity{
		repo: store,
		cfg:  cfg,
	}
}
