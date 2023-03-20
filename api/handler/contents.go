package handler

import (
	"fmt"
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

func (h *Handler) UpdateImageLink(c *gin.Context) {
	var req request.UploadImageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	imageLink, err := h.entity.UpdateImageLink(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"image_link": imageLink})
}

func (h *Handler) UpdateProgress(c *gin.Context) {
	var req request.UpdateProgressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(req)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	projectID, err := h.entity.UpdateProgress(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"project_id": projectID})
}
