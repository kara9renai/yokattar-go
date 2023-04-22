package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/kara9renai/yokattar-go/pkg/domain/object"
	"github.com/kara9renai/yokattar-go/pkg/http/middleware"
	"github.com/kara9renai/yokattar-go/pkg/server/handler/httperror"
	"github.com/kara9renai/yokattar-go/pkg/server/handler/request"
)

// Handle Request for `POST /v1/accounts/{username}/follow`
func (h *handler) Follow(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	relation := new(object.Relationship)
	username := request.UsernameOf(r)
	followingUser := middleware.AccountOf(r)

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
