package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kara9renai/yokattar-go/app/domain/object"
	"github.com/kara9renai/yokattar-go/app/domain/repository"
)

type (
	account struct {
		db *sqlx.DB
	}
)

func NewAccount(db *sqlx.DB) repository.Account {
	return &account{db: db}
}

func (r *account) FindByUsername(ctx context.Context, username string) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, "select * from account where username = ?", username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}

// CreateAccount: ユーザ名とパスワードからアカウントを新規作成する
func (r *account) CreateAccount(ctx context.Context, newAccount *object.Account) (*object.Account, error) {
	entity := new(object.Account)

	const (
		insert = `insert into account (
				username, password_hash, display_name, avatar, header, note
				)
				values (?, ?, ?, ?, ?, ?)`

		confirm = "select * from account where id = ?"
	)

	// prepared statement
	stmt, err := r.db.PreparexContext(ctx, insert)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(ctx,
		newAccount.Username,
		newAccount.PasswordHash,
		newAccount.DisplayName,
		newAccount.Avatar,
		newAccount.Header,
		newAccount.Note)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	err = r.db.QueryRowxContext(ctx, confirm, id).StructScan(entity)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}
