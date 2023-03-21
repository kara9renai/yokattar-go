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
	status struct {
		db *sqlx.DB
	}
)

func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

func (r *status) Create(ctx context.Context, accountId int64, content string) (*object.Status, error) {
	const (
		insert  = `insert into status (account_id, content) values (?, ?)`
		confirm = `select * from status where id = ?`
	)
	entity := new(object.Status)

	stmt, err := r.db.PreparexContext(ctx, insert)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, accountId, content)
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

func (r *status) Get(ctx context.Context, id int64) (*object.Status, error) {
	const (
		query = `select * from status where id = ?`
	)
	entity := new(object.Status)
	err := r.db.QueryRowxContext(ctx, query, id).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}

func (r *status) Delete(ctx context.Context, statusId int64) error {
	const (
		deleteFmt = `DELETE FROM status WHERE id = ?`
	)
	_, err := r.db.ExecContext(ctx, deleteFmt, statusId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return fmt.Errorf("%w", err)
	}

	return nil
}
