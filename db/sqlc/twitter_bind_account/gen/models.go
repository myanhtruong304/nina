// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import ()

type TwitterBindAccount struct {
	ID                int32  `json:"id"`
	ProjectName       string `json:"project_name"`
	AccessToken       string `json:"access_token"`
	AccessTokenSecret string `json:"access_token_secret"`
}