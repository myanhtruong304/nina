package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nina/api/handler"
	db "github.com/nina/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(r *gin.Engine, store db.Store, h *handler.Handler) *Server {
	v1 := r.Group("/api/v1")
	server := &Server{store: store}

	groupUsers := v1.Group("/users")
	{
		groupUsers.POST("/create-user", h.CreateUser)
	}

	groupProjects := v1.Group("/projects")
	{
		groupProjects.POST("/add-project", h.AddProject)
		groupProjects.POST("/upload-image", h.UploadImageToGCS)
	}

	server.router = r
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
