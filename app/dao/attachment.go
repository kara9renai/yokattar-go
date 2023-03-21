package dao

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kara9renai/yokattar-go/app/domain/object"
	"github.com/kara9renai/yokattar-go/app/domain/repository"
)

type (
	attachment struct {
		db *sqlx.DB
	}
)

func NewAttachment(db *sqlx.DB) repository.Attachment {
	return &attachment{db: db}
}

func (r *attachment) UploadFile(ctx context.Context, fileName string) (*object.Attachment, error) {
	const (
		insert  = `INSERT INTO attachment (type, url, description) VALUES (?, ?, ?)`
		confirm = `SELECT id, type, url, description FROM attachment WHERE id = ?`
	)

	entity := new(object.Attachment)
	result, err := r.db.ExecContext(ctx, insert, "image", fileName, "string")

	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()

	err = r.db.QueryRowxContext(ctx, confirm, id).StructScan(entity)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}
