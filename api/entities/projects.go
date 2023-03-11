package entities

import (
	"github.com/gin-gonic/gin"
	db "github.com/nina/db/sqlc"
)

func (e *Entity) AddProject(c *gin.Context, project db.AddProjectParams) (string, error) {
	projectName, err := e.repo.AddProject(c, project)

	if err != nil {
		return "", err
	}
	return projectName, nil
}
