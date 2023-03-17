package request

import (
	"time"

	"github.com/google/uuid"
)

type CreateContentRequest struct {
	ID          uuid.UUID
	Content     string
	ProjectName string
	ImageText   string
	CreatedAt   time.Time
}

type UpdateScheduleRequest struct {
	ScheduleTime time.Time
}

type UpdateProgressRequest struct {
	Facebook string
	Twitter  string
	Linkedin string
}

type UpdateImageLinkRequest struct {
	ImageLink   string
	LastUpdated time.Time
}
