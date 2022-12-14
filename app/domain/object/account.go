package object

import (
	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type (
	AccountID    = int64
	PasswordHash = string

	Account struct {
		ID AccountID `json:"-"`

		Username string `json:"username,omitempty"`

		PasswordHash string `json:"-" db:"password_hash"`

		DisplayName *string `json:"display_name,omitempty" db:"display_name"`

		Avatar *string `json:"avatar,omitempty"`

		Header *string `json:"header,omitempty"`

		Note *string `json:"note,omitempty"`

		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`

		// フォロワーの数
		FollowersCount int64 `json:"followers_count" db:"followers_count"`

		// フォローしているアカウントの数
		FollowingCount int64 `json:"following_count" db:"following_count"`
	}

	Relationship struct {
		ID AccountID `json:"id"`

		IsFollowing bool `json:"following"`

		IsFollowedby bool `json:"followed_by"`
	}
)

func (a *Account) CheckPassword(pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.PasswordHash), []byte(pass)) == nil
}

func (a *Account) SetPassword(pass string) error {
	passwordHash, err := generatePasswordHash(pass)
	if err != nil {
		return fmt.Errorf("generate error: %w", err)
	}
	a.PasswordHash = passwordHash
	return nil
}

func generatePasswordHash(pass string) (PasswordHash, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("hashing password failed: %w", errors.WithStack(err))
	}
	return PasswordHash(hash), nil
}
