package repository

import "context"

type Like interface {
	// 引数idのStatusをLIKEする
	LikeByStatusId(ctx context.Context, statusId int64) (int64, error)
	// LikeしているStatusを取得する（追加予定）
	// GetLikeStatus(ctx context.Context, acountId int64) ([]*object.Status, error)
}
