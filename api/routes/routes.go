package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nina/api/handler"
	db "github.com/nina/db/sqlc"
)

func NewRoute(r *gin.Engine, store db.Store, h *handler.Handler) gin.Engine {
	v1 := r.Group("/api/v1")

	groupUsers := v1.Group("/users")
	{
		groupUsers.POST("/create-user", h.CreateUser)
	}

	groupProjects := v1.Group("/projects")
	{
		groupProjects.POST("/add-project", h.AddProject)
	}

	groupContents := v1.Group("/contents")
	{
		groupContents.POST("/create-content", h.CreateContent)
		groupContents.POST("/update-image-link", h.UpdateImageLink)
		groupContents.POST("/update-progress", h.UpdateProgress)
	}

	return *r
}
