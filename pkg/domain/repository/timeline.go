package repository

import (
	"context"

	"github.com/kara9renai/yokattar-go/pkg/domain/object"
)

type Timeline interface {
	GetPublic(ctx context.Context, maxId int64, sinceId int64, limit int64) ([]*object.Status, error)
	GetHome(ctx context.Context, accountId int64, maxId, sinceId, limit int64) ([]*object.Status, error)
}
