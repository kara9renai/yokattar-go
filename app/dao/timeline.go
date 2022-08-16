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

func (r *timeline) GetPublicTimelines(ctx context.Context, maxId int64, sinceId int64, limit int64) ([]*object.Status, error) {

	var (
		entity []*object.Status
	)

	const (
		sql = `SELECT status.id, 
				status.account_id, 
				status.content, 
				status.create_at, 
				account.id "account.id", 
				account.username "account.username",
				account.display_name "account.display_name",
				account.create_at "account.create_at"
				FROM status
				LEFT JOIN account ON
				status.account_id = account.id
				LIMIT ?`
	)

	if maxId == 0 && sinceId == 0 {
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

	return entity, nil
}
