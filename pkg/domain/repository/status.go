package repository

import (
	"context"

	"github.com/kara9renai/yokattar-go/pkg/domain/object"
)

type Status interface {
	// create a status
	Create(ctx context.Context, accountId int64, status string) (*object.Status, error)
	// find a status by id
	Get(ctx context.Context, id int64) (*object.Status, error)
	// delete a status by id
	Delete(ctx context.Context, id int64) error
}
