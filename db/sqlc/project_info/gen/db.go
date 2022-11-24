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
	if q.createProjectInfoStmt, err = db.PrepareContext(ctx, createProjectInfo); err != nil {
		return nil, fmt.Errorf("error preparing query CreateProjectInfo: %w", err)
	}
	if q.deleteProjectStmt, err = db.PrepareContext(ctx, deleteProject); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteProject: %w", err)
	}
	if q.getListProjectsStmt, err = db.PrepareContext(ctx, getListProjects); err != nil {
		return nil, fmt.Errorf("error preparing query GetListProjects: %w", err)
	}
	if q.getOneProjectStmt, err = db.PrepareContext(ctx, getOneProject); err != nil {
		return nil, fmt.Errorf("error preparing query GetOneProject: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createProjectInfoStmt != nil {
		if cerr := q.createProjectInfoStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createProjectInfoStmt: %w", cerr)
		}
	}
	if q.deleteProjectStmt != nil {
		if cerr := q.deleteProjectStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteProjectStmt: %w", cerr)
		}
	}
	if q.getListProjectsStmt != nil {
		if cerr := q.getListProjectsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getListProjectsStmt: %w", cerr)
		}
	}
	if q.getOneProjectStmt != nil {
		if cerr := q.getOneProjectStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getOneProjectStmt: %w", cerr)
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
	createProjectInfoStmt *sql.Stmt
	deleteProjectStmt     *sql.Stmt
	getListProjectsStmt   *sql.Stmt
	getOneProjectStmt     *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                    tx,
		tx:                    tx,
		createProjectInfoStmt: q.createProjectInfoStmt,
		deleteProjectStmt:     q.deleteProjectStmt,
		getListProjectsStmt:   q.getListProjectsStmt,
		getOneProjectStmt:     q.getOneProjectStmt,
	}
}