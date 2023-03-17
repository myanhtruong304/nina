package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nina/api/request"
)

func (h *Handler) CreateContent(c *gin.Context) {
	var req request.CreateContentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	contentID, err := h.entity.CreateContent(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"content_id": contentID})
}
