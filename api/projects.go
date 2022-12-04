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

type updateProjectsParams struct {
	ProjectName     string `json:"project_name"`
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

func (server *Server) updateProjectInfo(ctx *gin.Context) {
	var req updateProjectsParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	projectInfo, err := server.store.GetOneProject(ctx, req.ProjectName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	arg := db.UpdateProjectInfoParams{
		ProjectName:     projectInfo.ProjectName,
		Symbol:          projectInfo.Symbol,
		ContractAddress: projectInfo.ContractAddress,
		Explorer:        projectInfo.Explorer,
		Twitter:         projectInfo.Twitter,
		Facebook:        projectInfo.Facebook,
		Linkedin:        projectInfo.Linkedin,
		Medium:          projectInfo.Medium,
		Telegram:        projectInfo.Telegram,
		Website:         projectInfo.Website,
		Git:             projectInfo.Git,
		Cmc:             projectInfo.Cmc,
		Coingecko:       projectInfo.Coingecko,
	}

	if req.Symbol != "" {
		arg.Symbol = sql.NullString{String: req.Symbol, Valid: true}
	}
	if req.ContractAddress != "" {
		arg.ContractAddress = sql.NullString{String: req.ContractAddress, Valid: true}
	}
	if req.Explorer != "" {
		arg.Explorer = sql.NullString{String: req.Explorer, Valid: true}
	}
	if req.Twitter != "" {
		arg.Twitter = sql.NullString{String: req.Twitter, Valid: true}
	}
	if req.Facebook != "" {
		arg.Facebook = sql.NullString{String: req.Facebook, Valid: true}
	}
	if req.Linkedin != "" {
		arg.Linkedin = sql.NullString{String: req.Linkedin, Valid: true}
	}
	if req.Medium != "" {
		arg.Medium = sql.NullString{String: req.Medium, Valid: true}
	}
	if req.Telegram != "" {
		arg.Telegram = sql.NullString{String: req.Telegram, Valid: true}
	}
	if req.Website != "" {
		arg.Website = sql.NullString{String: req.Website, Valid: true}
	}
	if req.Git != "" {
		arg.Git = sql.NullString{String: req.Git, Valid: true}
	}
	if req.Cmc != "" {
		arg.Cmc = sql.NullString{String: req.Cmc, Valid: true}
	}
	if req.Coingecko != "" {
		arg.Coingecko = sql.NullString{String: req.Coingecko, Valid: true}
	}

	project, err := server.store.UpdateProjectInfo(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
	}

	ctx.JSON(http.StatusOK, project)
}
