package db

import (
	content "github.com/nina/db/sqlc/content"
	imageContent "github.com/nina/db/sqlc/image_content"
	projectInfo "github.com/nina/db/sqlc/project_info"
	tags "github.com/nina/db/sqlc/tags"
	twitterBindAccount "github.com/nina/db/sqlc/twitter_bind_account"
)

type Store struct {
	ProjectInfo        projectInfo.Itf
	Content            content.Itf
	Tag                tags.Itf
	ImageContent       imageContent.Itf
	TwitterBindAccount twitterBindAccount.Itf
}
