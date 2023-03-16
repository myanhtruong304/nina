package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UploadImageToGCS(c *gin.Context) {
	var q struct {
		Path        string `json:"image" binding:"required"`
		ProjectName string `json:"project_name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url, err := h.entity.UploadImageToGCS(q.Path, q.ProjectName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"image_url": url})
}
