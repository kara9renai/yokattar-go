package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kara9renai/yokattar-go/app/domain/object"
	"github.com/kara9renai/yokattar-go/app/handler/auth"
	"github.com/kara9renai/yokattar-go/app/handler/httperror"
)

// Handle Request for `POST /accounts/{username}/unfollow`
func (h *handler) Unfollow(w http.ResponseWriter, r *http.Request) {

	relation := new(object.Relationship)

	ctx := r.Context()

	username := chi.URLParam(r, "username")

	unfollowingUser := auth.AccountOf(r)

	a := h.app.Dao.Account() // domain/repository の取得

	unfollowedUser, err := a.FindByUsername(ctx, username)

	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	if err := a.Unfollow(ctx, unfollowingUser.ID, unfollowedUser.ID); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
	// フォローを解除する対象が、同時にフォローしているのかどうかを確認する
	flag, err := a.FindRelationByID(ctx, unfollowedUser.ID, unfollowingUser.ID)

	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	relation.ID = unfollowingUser.ID
	relation.IsFollowing = false
	relation.IsFollowedby = flag

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(relation); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
