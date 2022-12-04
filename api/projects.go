package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/nina/db/sqlc"
)

type addProjectsParams struct {
	ProjectName     string `json:"project_name" binding:"required"`
	Symbol          string `json:"symbol"`
	ContractAddress string `json:"contract_address"`
	Explorer        string `json:"explorer"`
	Twitter         string `json:"twitter"`
	Facebook        string `json:"facebook"`
	Linkedin        string `json:"linkedin"`
	Medium          string `json:"medium"`
	Telegram        string `json:"telegram"`
	Website         string `json:"website"`
	Git             string `json:"git"`
	Cmc             string `json:"cmc"`
	Coingecko       string `json:"coingecko"`
}

func (server *Server) addProjects(ctx *gin.Context) {
	var req addProjectsParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.CreateProjectInfoParams{
		ProjectName:     req.ProjectName,
		Symbol:          sql.NullString{String: req.Symbol, Valid: true},
		ContractAddress: sql.NullString{String: req.ContractAddress, Valid: true},
		Explorer:        sql.NullString{String: req.Explorer, Valid: true},
		Twitter:         sql.NullString{String: req.Twitter, Valid: true},
		Facebook:        sql.NullString{String: req.Facebook, Valid: true},
		Linkedin:        sql.NullString{String: req.Linkedin, Valid: true},
		Medium:          sql.NullString{String: req.Medium, Valid: true},
		Telegram:        sql.NullString{String: req.Telegram, Valid: true},
		Website:         sql.NullString{String: req.Website, Valid: true},
		Git:             sql.NullString{String: req.Website, Valid: true},
		Cmc:             sql.NullString{String: req.Cmc, Valid: true},
		Coingecko:       sql.NullString{String: req.Coingecko, Valid: true},
	}

	project, err := server.store.CreateProjectInfo(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
	}

	ctx.JSON(http.StatusOK, project)
}
