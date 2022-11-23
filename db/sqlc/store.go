package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// Store provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (s *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rollback err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type CreateImageTx struct {
	ContentID int       `json:"content_id"`
	ImageLink string    `json:"image_link"`
	UpdatedAt time.Time `json:"updated_time"`
}

type UpdateImageTxResult struct {
	Image   ImageContent `json:"image"`
	Content Content      `json:"content"`
}

// Assume that image_content existed.
// update image_link + updated_time into image_content table
// update image_link + image_id + updated_time to content table
func (s *Store) UpdateImageLinkTx(ctx context.Context, arg CreateImageTx) (UpdateImageTxResult, error) {
	var result UpdateImageTxResult

	err := s.execTx(ctx, func(q *Queries) error {
		var err error

		result.Image, err = q.UpdateLink(ctx, UpdateLinkParams{
			ContentID: int32(arg.ContentID),
			Link:      sql.NullString{String: arg.ImageLink, Valid: true},
			UpdatedAt: arg.UpdatedAt,
		})
		if err != nil {
			return err
		}

		result.Content, err = q.UpdateImageLink(ctx, UpdateImageLinkParams{
			ID:        int32(arg.ContentID),
			ImageLink: sql.NullString{String: arg.ImageLink, Valid: true},
			UpdatedAt: arg.UpdatedAt,
		})
		if err != nil {
			return err
		}
		return nil
	})

	return result, err
}
