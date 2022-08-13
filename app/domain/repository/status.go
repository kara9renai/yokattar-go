package repository

import (
	"context"

	"github.com/kara9renai/yokattar-go/app/domain/object"
)

type Status interface {
	CreateStatus(ctx context.Context, accountId int64, status string) (*object.Status, error)
}
