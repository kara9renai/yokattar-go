package repository

import (
	"context"

	"github.com/kara9renai/yokattar-go/app/domain/object"
)

type Attachment interface {
	Save(ctx context.Context, fileName string) (*object.Attachment, error)
}
