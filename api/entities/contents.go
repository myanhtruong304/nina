package entities

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nina/api/request"
	db "github.com/nina/db/sqlc"
	"github.com/nina/util"
)

func (e *Entity) CreateContent(c *gin.Context, req request.CreateContentRequest) (string, error) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	project, err := e.repo.GetProject(c, req.ProjectName)
	if err != nil {
		return "", err
	}

	ID := fmt.Sprintf("%v_%v", project.Symbol.String, timestamp)

	wordCount := int32(len(req.Content))
	contentType := "Tweet"
	if !util.ContainsLetters(req.Content) {
		contentType = "Retweet"
	}

	if req.ImageText != "" {
		contentType = "Image tweet"
	}

	content := db.AddContentParams{
		ID:          ID,
		Content:     req.Content,
		WordCount:   wordCount,
		ContentType: sql.NullString{String: contentType, Valid: true},
		ProjectName: req.ProjectName,
		ImageText:   sql.NullString{String: req.ImageText, Valid: true},
		CreatedAt:   time.Now(),
	}
	contentID, err := e.repo.AddContent(c, content)
	if err != nil {
		return "", err
	}

	return contentID, nil
}

func (e *Entity) UpdateImageLink(c *gin.Context, req request.UploadImageRequest) (string, error) {
	imageGcsUrl, err := e.UploadImageToGCS(c, req.FilePath, req.ID)
	if err != nil {
		return "", err
	}

	updateParam := db.UpdateImageLinkParams{
		ID:            req.ID,
		ImageLink:     sql.NullString{String: imageGcsUrl, Valid: true},
		LastUpdatedAt: time.Now(),
	}

	imageLink, err := e.repo.UpdateImageLink(c, updateParam)
	if err != nil {
		return "", err
	}

	return imageLink.String, nil
}

func (e *Entity) UpdateProgress(c *gin.Context, req request.UpdateProgressRequest) (string, error) {
	updateParam := db.UpdateProgressParams{
		ID:            req.ContentID,
		FacebookCheck: sql.NullString{String: req.Facebook, Valid: true},
		TwitterCheck:  sql.NullString{String: req.Twitter, Valid: true},
		LinkedinCheck: sql.NullString{String: req.Linkedin, Valid: true},
	}

	contentID, err := e.repo.UpdateProgress(c, updateParam)
	if err != nil {
		return "", err
	}

	return contentID, nil
}
