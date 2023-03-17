// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.addContentStmt, err = db.PrepareContext(ctx, addContent); err != nil {
		return nil, fmt.Errorf("error preparing query AddContent: %w", err)
	}
	if q.addProjectStmt, err = db.PrepareContext(ctx, addProject); err != nil {
		return nil, fmt.Errorf("error preparing query AddProject: %w", err)
	}
	if q.addUserStmt, err = db.PrepareContext(ctx, addUser); err != nil {
		return nil, fmt.Errorf("error preparing query AddUser: %w", err)
	}
	if q.getAllProjectStmt, err = db.PrepareContext(ctx, getAllProject); err != nil {
		return nil, fmt.Errorf("error preparing query GetAllProject: %w", err)
	}
	if q.getProjectStmt, err = db.PrepareContext(ctx, getProject); err != nil {
		return nil, fmt.Errorf("error preparing query GetProject: %w", err)
	}
	if q.modifyProjectInfoStmt, err = db.PrepareContext(ctx, modifyProjectInfo); err != nil {
		return nil, fmt.Errorf("error preparing query ModifyProjectInfo: %w", err)
	}
	if q.updateImageLinkStmt, err = db.PrepareContext(ctx, updateImageLink); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateImageLink: %w", err)
	}
	if q.updateProgressStmt, err = db.PrepareContext(ctx, updateProgress); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateProgress: %w", err)
	}
	if q.updateScheduleStmt, err = db.PrepareContext(ctx, updateSchedule); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateSchedule: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.addContentStmt != nil {
		if cerr := q.addContentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addContentStmt: %w", cerr)
		}
	}
	if q.addProjectStmt != nil {
		if cerr := q.addProjectStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addProjectStmt: %w", cerr)
		}
	}
	if q.addUserStmt != nil {
		if cerr := q.addUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing addUserStmt: %w", cerr)
		}
	}
	if q.getAllProjectStmt != nil {
		if cerr := q.getAllProjectStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAllProjectStmt: %w", cerr)
		}
	}
	if q.getProjectStmt != nil {
		if cerr := q.getProjectStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProjectStmt: %w", cerr)
		}
	}
	if q.modifyProjectInfoStmt != nil {
		if cerr := q.modifyProjectInfoStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing modifyProjectInfoStmt: %w", cerr)
		}
	}
	if q.updateImageLinkStmt != nil {
		if cerr := q.updateImageLinkStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateImageLinkStmt: %w", cerr)
		}
	}
	if q.updateProgressStmt != nil {
		if cerr := q.updateProgressStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateProgressStmt: %w", cerr)
		}
	}
	if q.updateScheduleStmt != nil {
		if cerr := q.updateScheduleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateScheduleStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                    DBTX
	tx                    *sql.Tx
	addContentStmt        *sql.Stmt
	addProjectStmt        *sql.Stmt
	addUserStmt           *sql.Stmt
	getAllProjectStmt     *sql.Stmt
	getProjectStmt        *sql.Stmt
	modifyProjectInfoStmt *sql.Stmt
	updateImageLinkStmt   *sql.Stmt
	updateProgressStmt    *sql.Stmt
	updateScheduleStmt    *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                    tx,
		tx:                    tx,
		addContentStmt:        q.addContentStmt,
		addProjectStmt:        q.addProjectStmt,
		addUserStmt:           q.addUserStmt,
		getAllProjectStmt:     q.getAllProjectStmt,
		getProjectStmt:        q.getProjectStmt,
		modifyProjectInfoStmt: q.modifyProjectInfoStmt,
		updateImageLinkStmt:   q.updateImageLinkStmt,
		updateProgressStmt:    q.updateProgressStmt,
		updateScheduleStmt:    q.updateScheduleStmt,
	}
}
