package dao

import (
	"context"

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
	// impl business logic
	return nil, nil
}
