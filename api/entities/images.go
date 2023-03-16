package entities

import (
	"context"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"cloud.google.com/go/storage"
	"github.com/k0kubun/pp"
	"google.golang.org/api/option"
)

func (e *Entity) UploadImageToGCS(path, project string) (string, error) {
	APIKey := e.cfg.GoogleServiceAPIKey

	pp.Println(APIKey)
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

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	bucket := client.Bucket("test_myanh")
	wc := bucket.Object(project + "_" + timestamp).NewWriter(ctx)

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

	return fmt.Sprintf("https://storage.googleapis.com/test_myanh/%s_%s", project, timestamp), nil

}
