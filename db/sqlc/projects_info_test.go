package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateProject(t *testing.T) {
	arg := CreateProjectInfoParams{
		ProjectName: "Supreme Finance",
		Twitter:     sql.NullString{String: "https://twitter.com/SupremeFinance2", Valid: true},
		Facebook:    sql.NullString{String: "https://www.facebook.com/Supreme-Finance-100106618907034", Valid: true},
		Linkedin:    sql.NullString{String: "https://www.linkedin.com/company/supreme-finance", Valid: true},
		Medium:      sql.NullString{String: "https://supremefinance.medium.com/", Valid: true},
		Telegram:    sql.NullString{String: "https://t.me/SUPREMEFINANCEENGLISH", Valid: true},
		Website:     sql.NullString{String: "https://www.supremefinance.io/", Valid: true},
		Git:         sql.NullString{},
		Cmc:         sql.NullString{String: "https://coinmarketcap.com/currencies/hype/", Valid: true},
		Coingecko:   sql.NullString{String: "https://www.coingecko.com/en/coins/supreme-finance", Valid: true},
	}

	project, err := testQueries.CreateProjectInfo(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, project)

	require.Equal(t, arg.ProjectName, project.ProjectName)
	require.NotZero(t, project.CreatedAt)
}
func TestGetProject(t *testing.T) {
	project, err := testQueries.GetOneProject(context.Background(), "Supreme Finance")
	t.Log(project.Coingecko.String)
	require.NoError(t, err)
	require.NotEmpty(t, project)
}

func TestGetListProject(t *testing.T) {
	project, err := testQueries.GetListProjects(context.Background())
	t.Log(project[0].Coingecko.String)
	require.NoError(t, err)
	require.NotEmpty(t, project)
}

// func TestDeleteProject(t *testing.T) {
// 	err := testQueries.DeleteProject(context.Background(), "Supreme Finance")
// 	require.NoError(t, err)
// }
