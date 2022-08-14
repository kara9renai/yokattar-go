package repository

import (
	"context"

	"github.com/kara9renai/yokattar-go/app/domain/object"
)

type Account interface {
	FindByUsername(ctx context.Context, username string) (*object.Account, error)
	CreateAccount(ctx context.Context, user *object.Account) (*object.Account, error)
	FindByID(ctx context.Context, accoundId int64) (*object.Account, error)
}
