package entities

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func (e *Entity) UploadImageToGCS(c *gin.Context, path, content_ID string) (string, error) {
	APIKey := e.cfg.GoogleServiceAPIKey
	project := strings.Split(content_ID, "_")[0]

	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithAPIKey(APIKey))
	if err != nil {
		return "", err
	}
	defer client.Close()

	// open the local image file
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return "", err
	}
	defer file.Close()

	bucket := client.Bucket("test_myanh")
	wc := bucket.Object(project + "/" + content_ID).NewWriter(ctx)

	_, err = io.Copy(wc, file)
	if err != nil {
		fmt.Printf("Failed to write file to GCS: %v\n", err)
		return "", err
	}

	err = wc.Close()
	if err != nil {
		fmt.Printf("Failed to close GCS writer: %v\n", err)
		return "", err
	}

	return fmt.Sprintf("https://storage.googleapis.com/test_myanh/%s/%s", project, content_ID), nil

}
