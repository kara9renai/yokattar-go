package repository

import (
	"context"

	"github.com/kara9renai/yokattar-go/app/domain/object"
)

type Account interface {
	FindByUsername(ctx context.Context, username string) (*object.Account, error)
	CreateAccount(ctx context.Context, user *object.Account) (*object.Account, error)
	FindByID(ctx context.Context, accoundId int64) (*object.Account, error)
	// impl function to follow
	Follow(ctx context.Context, followingId int64, followerId int64) error
	// impl function to get relationship
	FindRelationByID(ctx context.Context, followingId int64, followerId int64) (bool, error)
	// impl function to get  who account is following
	FindFollowing(ctx context.Context, accountId int64, limit int64) ([]*object.Account, error)
	// impl function to get who account's followers
	FindFollowers(ctx context.Context, accountId int64, limit int64) ([]*object.Account, error)
}
