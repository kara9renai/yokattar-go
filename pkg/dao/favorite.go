package dao

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/kara9renai/yokattar-go/pkg/domain/object"
	"github.com/kara9renai/yokattar-go/pkg/domain/repository"
)

type (
	favorite struct {
		db *sqlx.DB
	}
)

func NewFavorite(db *sqlx.DB) repository.Favorite {
	return &favorite{db: db}
}

func (r *favorite) Create(ctx context.Context, accountId int64, statusId int64) error {
	const (
		insert = `INSERT INTO favorite ( account_id, status_id ) VALUES (?, ?)`
		update = `UPDATE status SET favorite_coun = favorite_coun + 1 where id = ?`
	)
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, insert, accountId, statusId)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.ExecContext(ctx, update, statusId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
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
		sql = `SELECT id, create_at FROM favorite WHERE account_id = ? AND status_id = ?`
	)
	entity := new(object.Favorite)
	err := r.db.QueryRowxContext(ctx, sql, accountId, statusId).StructScan(entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *favorite) Delete(ctx context.Context, accountId int64, statusId int64) error {
	const sql = `DELETE FROM favorite WHERE account_id = ? and status_id = ?`
	_, err := r.db.ExecContext(ctx, sql, accountId, statusId)
	if err != nil {
		return err
	}
	return nil
}
