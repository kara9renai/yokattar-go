package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/kara9renai/yokattar-go/pkg/domain/object"
	"github.com/kara9renai/yokattar-go/pkg/http/middleware"
	"github.com/kara9renai/yokattar-go/pkg/server/handler/httperror"
	"github.com/kara9renai/yokattar-go/pkg/server/handler/request"
)

// Handle Request for `POST /accounts/{username}/unfollow`
func (h *handler) Unfollow(w http.ResponseWriter, r *http.Request) {
	relation := new(object.Relationship)
	ctx := r.Context()
	username := request.UsernameOf(r)
	unfollowingUser := middleware.AccountOf(r)

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
