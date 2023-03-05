package dao

import (
	"context"

	"github.com/jmoiron/sqlx"
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

func (r *like) LikeByStatusId(ctx context.Context, statusId int64) (int64, error) {
	return 0, nil
}
