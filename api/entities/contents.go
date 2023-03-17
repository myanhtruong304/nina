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
	ID := fmt.Sprintf("%v_%v", req.ProjectName, timestamp)

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
