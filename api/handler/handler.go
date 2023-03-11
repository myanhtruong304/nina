package handler

import (
	"github.com/nina/api/entities"
	"github.com/nina/util"
)

type Handler struct {
	entity *entities.Entity
}

func NewHandler(cfg *util.Config) *Handler {
	return &Handler{
		entity: entities.NewEntity(cfg),
	}
}
