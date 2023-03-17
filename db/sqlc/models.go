// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"
	"time"
)

type Content struct {
	ID            string         `json:"id"`
	Content       string         `json:"content"`
	WordCount     string         `json:"word_count"`
	ScheduleTime  sql.NullString `json:"schedule_time"`
	FacebookCheck sql.NullString `json:"facebook_check"`
	TwitterCheck  sql.NullString `json:"twitter_check"`
	LinkedinCheck sql.NullString `json:"linkedin_check"`
	ContentType   sql.NullString `json:"content_type"`
	ProjectName   string         `json:"project_name"`
	ImageText     sql.NullString `json:"image_text"`
	ImageLink     sql.NullString `json:"image_link"`
	CreatedAt     time.Time      `json:"created_at"`
	LastUpdatedAt time.Time      `json:"last_updated_at"`
}

type Project struct {
	ProjectName     string         `json:"project_name"`
	Symbol          sql.NullString `json:"symbol"`
	ContractAddress sql.NullString `json:"contract_address"`
	Explorer        sql.NullString `json:"explorer"`
	Website         sql.NullString `json:"website"`
	Twitter         sql.NullString `json:"twitter"`
	Facebook        sql.NullString `json:"facebook"`
	Linkedin        sql.NullString `json:"linkedin"`
	Telegram        sql.NullString `json:"telegram"`
	Medium          sql.NullString `json:"medium"`
	Whitepaper      sql.NullString `json:"whitepaper"`
	Email           sql.NullString `json:"email"`
	Owner           string         `json:"owner"`
	CreatedAt       time.Time      `json:"created_at"`
}

type User struct {
	Username     string    `json:"username"`
	HashedPwd    string    `json:"hashed_pwd"`
	UserEmail    string    `json:"user_email"`
	PwdChangedAt time.Time `json:"pwd_changed_at"`
	CreatedAt    time.Time `json:"created_at"`
}
