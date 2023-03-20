package request

import (
	"time"
)

type CreateContentRequest struct {
	Content     string `json:"content"`
	ProjectName string `json:"project_name"`
	ImageText   string `json:"image_text"`
}

type UpdateScheduleRequest struct {
	ScheduleTime time.Time `json:"schedule_time"`
}

type UpdateProgressRequest struct {
	ContentID string `json:"id" bind:"require"`
	Facebook  string `json:"facebook"`
	Twitter   string `json:"twitter"`
	Linkedin  string `json:"linkedin"`
}

type UploadImageRequest struct {
	ID       string `json:"id"`
	FilePath string `json:"image_path"`
}
