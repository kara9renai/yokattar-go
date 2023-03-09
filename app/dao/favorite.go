package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

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

func (r *favorite) FavoriteByStatusId(ctx context.Context, accountId int64, statusId int64) (*object.Favorite, error) {
	const (
		insert  = `INSERT INTO favorite ( account_id, status_id ) VALUES (?, ?)`
		confirm = `SELECT id, create_at FROM favorite WHERE id = ?`
	)
	entity := new(object.Favorite)

	stmt, err := r.db.PreparexContext(ctx, insert)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, accountId, statusId)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	err = r.db.QueryRowxContext(ctx, confirm, id).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("%w", err)
	}
	return entity, nil
}
