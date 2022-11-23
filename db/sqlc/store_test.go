package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestUpdateContent(t *testing.T) {
	store := NewStore(testDB)

	n := 5
	link := sql.NullString{String: "https://drive.google.com/file/d/1ruqbOMCXcovbT9Z_hLwdl53PUIQgd8QC/view?usp=sharing", Valid: true}

	errs := make(chan error)
	results := make(chan UpdateImageTxResult)
	for i := 0; i < n; i++ {
		go func() {
			result, err := store.UpdateImageLinkTx(context.Background(), CreateImageTx{
				ContentID: 1,
				ImageLink: link.String,
				UpdatedAt: time.Now(),
			})

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)
		require.Equal(t, result.Content.ImageLink, link)

		_, err = store.GetContentOneProject(context.Background(), "Supreme Finance")
		require.NoError(t, err)
	}

}
