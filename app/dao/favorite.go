package dao

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/kara9renai/yokattar-go/app/domain/object"
	"github.com/kara9renai/yokattar-go/app/domain/repository"
)

type (
	favorite struct {
		db *sqlx.DB
	}
)

func NewFavorite(db *sqlx.DB) repository.Favorite {
	return &favorite{db: db}
}

func (r *favorite) Favorite(ctx context.Context, accountId int64, statusId int64) (bool, error) {
	const (
		insert = `INSERT INTO favorite ( account_id, status_id ) VALUES (?, ?)`
	)
	stmt, err := r.db.PreparexContext(ctx, insert)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, accountId, statusId)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *favorite) Confirm(ctx context.Context, accountId int64, statusId int64) (bool, error) {
	const (
		sql = `SELECT * FROM favorite WHERE account_id = ? AND status_id = ?`
	)
	rows, err := r.db.QueryxContext(ctx, sql, accountId, statusId)
	if err != nil {
		return false, err
	}
	if rows.Next() {
		return true, nil
	} else {
		return false, nil
	}
}

func (r *favorite) Get(ctx context.Context, accountId int64, statusId int64) (*object.Favorite, error) {
	const (
		sql = `SELECT id, create_at FROM WHERE account_id = ? AND status_id = ?`
	)
	entity := new(object.Favorite)
	err := r.db.QueryRowxContext(ctx, sql, accountId, statusId).StructScan(entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}
