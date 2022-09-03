package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/kara9renai/yokattar-go/app/domain/object"
	"github.com/kara9renai/yokattar-go/app/handler/auth"
	"github.com/kara9renai/yokattar-go/app/handler/httperror"
	"github.com/kara9renai/yokattar-go/app/handler/request"
)

// Handle Request for `POST /v1/accounts/{username}/follow`
func (h *handler) Follow(w http.ResponseWriter, r *http.Request) {

	relation := new(object.Relationship)

	ctx := r.Context()

	username := request.UsernameOf(r)

	followingUser := auth.AccountOf(r)

	a := h.app.Dao.Account() // domain/repository の取得

	followedUser, err := a.FindByUsername(ctx, username)

	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	if err := a.Follow(ctx, followingUser.ID, followedUser.ID); err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	// フォローする対象が、同時に自分をフォローしているかどうかを確認する
	flag, err := a.FindRelationByID(ctx, followedUser.ID, followingUser.ID)

	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	relation.ID = followingUser.ID
	relation.IsFollowing = true
	relation.IsFollowedby = flag

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(relation); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
