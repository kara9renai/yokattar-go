package repository

import (
	"context"

	"github.com/kara9renai/yokattar-go/app/domain/object"
)

type Like interface {
	// 引数idのStatusをLIKEする
	LikeByStatusId(ctx context.Context, accountId int64, statusId int64) (*object.Like, error)
	// LikeしているStatusを取得する（追加予定）
	// GetLikeStatus(ctx context.Context, acountId int64) ([]*object.Status, error)
}
