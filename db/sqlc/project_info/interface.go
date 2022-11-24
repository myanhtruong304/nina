package db

import (
	"context"

	db "github.com/nina/db/sqlc/project_info/gen"
)

type Itf interface {
	CreateProjectInfo(ctx context.Context, arg db.CreateProjectInfoParams) (db.ProjectsInfo, error)
	DeleteProject(ctx context.Context, projectName string) error
	GetListProjects(ctx context.Context) ([]db.ProjectsInfo, error)
	GetOneProject(ctx context.Context, projectName string) (db.ProjectsInfo, error)
}
