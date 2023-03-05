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
	like struct {
		db *sqlx.DB
	}
)

func NewLike(db *sqlx.DB) repository.Like {
	return &like{db: db}
}

func (r *like) LikeByStatusId(ctx context.Context, accountId int64, statusId int64) (*object.Like, error) {
	const (
		insert  = `INSERT INTO favorite ( account_id, status_id ) VALUES (?, ?)`
		confirm = `SELECT id, create_at FROM favorite WHERE id = ?`
	)
	entity := new(object.Like)

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
