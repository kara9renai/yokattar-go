package repository

import (
	"context"

	"github.com/kara9renai/yokattar-go/app/domain/object"
)

type Status interface {
	// create a status
	CreateStatus(ctx context.Context, accountId int64, status string) (*object.Status, error)
	// find a status by id
	GetStatus(ctx context.Context, id int64) (*object.Status, error)
	// delete a status by id
	DeleteStatus(ctx context.Context, id int64) error
}
