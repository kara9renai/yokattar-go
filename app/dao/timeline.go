package dao

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kara9renai/yokattar-go/app/domain/object"
	"github.com/kara9renai/yokattar-go/app/domain/repository"
)

type (
	timeline struct {
		db *sqlx.DB
	}
)

func NewTimeline(db *sqlx.DB) repository.Timeline {
	return &timeline{db: db}
}

func (r *timeline) GetPublic(ctx context.Context, maxId int64, sinceId int64, limit int64) ([]*object.Status, error) {

	var (
		entity []*object.Status

		sql string = `SELECT status.id, 
				status.account_id, 
				status.content, 
				status.create_at, 
				account.id "account.id", 
				account.username "account.username",
				account.display_name "account.display_name",
				account.create_at "account.create_at"
				FROM status
				LEFT JOIN account ON
				status.account_id = account.id`
	)

	// max_id, since_id が指定されない場合
	// limit (default = 40) まで表示
	if maxId == 0 && sinceId == 0 {

		sql += ` LIMIT ?`
		rows, err := r.db.QueryxContext(ctx, sql, limit)

		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		for rows.Next() {
			var s object.Status
			err = rows.StructScan(&s)

			if err != nil {
				return nil, fmt.Errorf("%w", err)
			}

			entity = append(entity, &s)
		}

		return entity, nil
	}

	if maxId != 0 && sinceId != 0 {

		sql += ` WHERE status.id BETWEEN ? AND ? LIMIT ? `

		rows, err := r.db.QueryxContext(ctx, sql, sinceId, maxId, limit)

		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		for rows.Next() {
			var s object.Status
			err = rows.StructScan(&s)

			if err != nil {
				return nil, fmt.Errorf("%w", err)
			}

			entity = append(entity, &s)

		}

		return entity, nil
	}

	return entity, nil
}

func (r *timeline) GetHome(ctx context.Context, accountId int64, limit int64) ([]*object.Status, error) {

	var (
		entity []*object.Status

		sql string = `SELECT status.id, 
				status.account_id, 
				status.content, 
				status.create_at, 
				account.id "account.id", 
				account.username "account.username",
				account.display_name "account.display_name",
				account.create_at "account.create_at",
				account.followers_count "account.followers_count",
				account.following_count "account.following_count",
				account.note "account.note",
				account.avatar "account.avatar",
				account.header "account.header"
				FROM status
				INNER JOIN relation ON
				status.account_id = relation.followee_id
				LEFT JOIN account ON
				status.account_id = account.id
				WHERE relation.follower_id = ?
				LIMIT ?`
	)

	rows, err := r.db.QueryxContext(ctx, sql, accountId, limit)

	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	for rows.Next() {
		var s object.Status
		err = rows.StructScan(&s)

		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		entity = append(entity, &s)

	}

	return entity, nil

}
