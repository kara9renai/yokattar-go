package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/kara9renai/yokattar-go/pkg/domain/object"
	"github.com/kara9renai/yokattar-go/pkg/domain/repository"
	"github.com/kara9renai/yokattar-go/pkg/dto"
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
	const (
		confirm = `select * from account where username = ?`
	)
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, confirm, username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}

// CreateAccount: ユーザ名とパスワードからアカウントを新規作成する
func (r *account) Create(ctx context.Context, newAccount *object.Account) (*object.Account, error) {
	const (
		insert = `insert into account (
				username, password_hash, display_name, avatar, header, note,
				following_count, followers_count
				)
				values (?, ?, ?, ?, ?, ?, ?, ?)`

		confirm = "select * from account where id = ?"
	)
	entity := new(object.Account)

	// prepared statement
	stmt, err := r.db.PreparexContext(ctx, insert)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
		newAccount.Username,
		newAccount.PasswordHash,
		newAccount.DisplayName,
		newAccount.Avatar,
		newAccount.Header,
		newAccount.Note,
		0,
		0,
	)
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

func (r *account) FindByID(ctx context.Context, accountId int64) (*object.Account, error) {

	const (
		query = `SELECT * FROM account where id = ?`
	)

	entity := new(object.Account)

	err := r.db.QueryRowxContext(ctx, query, accountId).StructScan(entity)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}

// Follow: followerIdとfolloweeIdからフォロー関係を記録する
// TODO: ここはトランザクションで書きたい
func (r *account) Follow(ctx context.Context, followerId int64, followeeId int64) error {

	const (
		insert          = `INSERT INTO relation (follower_id, followee_id) VALUES (?, ?)`
		followingUpdate = `UPDATE account a SET following_count = following_count + 1 WHERE id = ?`
		followersUpdate = `UPDATE account a SET followers_count = followers_count + 1 WHERE id = ?`
	)

	tx := r.db.MustBeginTx(ctx, &sql.TxOptions{})
	defer tx.Rollback()

	tx.MustExecContext(ctx, insert, followerId, followeeId)

	tx.MustExecContext(ctx, followingUpdate, followerId)

	tx.MustExecContext(ctx, followersUpdate, followeeId)

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// FindRelationByID: followerId(フォローする側のID)とfolloweeId(フォローされる側)から
// 該当するリレーションを見つける関数
func (r *account) FindRelationByID(ctx context.Context, followerId int64, followeeId int64) (bool, error) {

	const (
		query = `SELECT * FROM relation WHERE follower_id = ? AND followee_id = ?`
	)

	rows, err := r.db.QueryxContext(ctx, query, followerId, followeeId)

	if err != nil {
		return false, nil
	}

	if rows.Next() {
		return true, nil
	} else {
		return false, nil
	}
}

// FindFollowing: パラメータで渡されたaccountがフォローしているaccountを返す関数
func (r *account) FindFollowing(ctx context.Context, accountId int64, limit int64) ([]*object.Account, error) {

	var entity []*object.Account

	const sql = `SELECT a.* FROM account a 
				INNER JOIN relation r
				ON a.id = r.followee_id
				WHERE r.follower_id = ?
				LIMIT ?`

	rows, err := r.db.QueryxContext(ctx, sql, accountId, limit)

	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	for rows.Next() {
		var a object.Account
		err = rows.StructScan(&a)

		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		entity = append(entity, &a)
	}

	return entity, nil
}

func (r *account) FindFollowers(ctx context.Context, accountId int64, limit int64) ([]*object.Account, error) {

	var entity []*object.Account

	const sql = `SELECT a.* FROM account a
				INNER JOIN relation r
				ON a.id = r.follower_id
				WHERE r.followee_id = ?
				LIMIT ?`

	rows, err := r.db.QueryxContext(ctx, sql, accountId, limit)

	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	for rows.Next() {
		var a object.Account
		err = rows.StructScan(&a)

		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		entity = append(entity, &a)
	}

	return entity, nil
}

// Unfollow: POSTをリクエストしたfollower(フォローする側)が、パラメータで指定されたfollowee(フォローされている側)の
// フォローを解除する関数
func (r *account) Unfollow(ctx context.Context, followerId int64, followeeId int64) error {

	const (
		deleteFmt       = `DELETE FROM relation WHERE follower_id = ? AND followee_id = ?`
		followingUpdate = `UPDATE account SET following_count = following_count - 1 WHERE id = ?`
		followersUpdate = `UPDATE account SET followers_count = followers_count - 1 WHERE id = ?`
	)

	result, err := r.db.ExecContext(ctx, deleteFmt, followerId, followeeId)

	if err != nil {
		return err
	}

	log.Println(result.RowsAffected())

	if affected, _ := result.RowsAffected(); affected == 0 {
		return errors.New("1行も影響を受けていません")
	}

	_, err = r.db.ExecContext(ctx, followingUpdate, followerId)

	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, followersUpdate, followeeId)

	if err != nil {
		return err
	}

	return nil
}

func (r *account) UpdateCredentials(ctx context.Context, userId int64, dto dto.Credentials) (*object.Account, error) {
	const (
		confirm = `SELECT * FROM account WHERE id = ?`
	)
	update := `UPDATE account SET `
	entity := new(object.Account)
	args := make([]interface{}, 0, 5)

	m := map[string]string{
		"display_name": dto.DisplayName,
		"avatar":       dto.Avatar,
		"header":       dto.Header,
		"note":         dto.Note,
	}
	isColumn := false

	for columnName, dtoValue := range m {
		if len(dtoValue) != 0 {
			if isColumn {
				update += " ,"
			}
			update += fmt.Sprintf(" %v = ? ", columnName)
			args = append(args, dtoValue)
			isColumn = true
		}
	}
	update += " WHERE id = ?"
	args = append(args, userId)

	stmt, err := r.db.PreparexContext(ctx, update)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		return nil, err
	}

	if err := r.db.QueryRowxContext(ctx, confirm, userId).StructScan(entity); err != nil {
		return nil, err
	}

	return entity, nil
}
