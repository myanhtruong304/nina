package entities

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nina/api/request"
	db "github.com/nina/db/sqlc"
	"github.com/nina/util"
)

func (e *Entity) CreateUser(c *gin.Context, user *request.CreateUser) (*db.AddUserRow, error) {
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	arg := db.AddUserParams{
		Username:     user.UserName,
		HashedPwd:    hashedPassword,
		UserEmail:    user.Email,
		PwdChangedAt: time.Now(),
		CreatedAt:    time.Now(),
	}
	data, err := e.repo.AddUser(c, arg)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
