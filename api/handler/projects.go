package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nina/api/request"
	db "github.com/nina/db/sqlc"
)

func (h *Handler) AddProject(c *gin.Context) {
	var req request.AddProject
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	arg := db.AddProjectParams{
		ProjectName:     req.ProjectName,
		Symbol:          req.Symbol,
		ContractAddress: req.ContractAddress,
		Owner:           req.Owner,
		CreatedAt:       time.Now(),
	}
	projectName, err := h.entity.AddProject(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error 1": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("project %v has been added.", projectName)})
}
