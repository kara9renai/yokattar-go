package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kara9renai/yokattar-go/app/config"
	"github.com/kara9renai/yokattar-go/app/handler/httperror"
	"github.com/kara9renai/yokattar-go/app/handler/request"
)

// Handle Request for `GET /accounts/{username}/following`
func (h *handler) Following(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	limit, err := request.URLParamOf(r, "limit")

	if err != nil {
		limit = config.DEFAULT_LIMIT
	}

	if limit > config.MAX_LIMIT {
		limit = config.MAX_LIMIT
	}

	username := chi.URLParam(r, "username")

	a := h.app.Dao.Account() // domain/repository の取得

	account, err := a.FindByUsername(ctx, username)

	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	followingUsers, err := a.FindFollowing(ctx, account.ID, limit)

	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(followingUsers); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
