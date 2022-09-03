package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kara9renai/yokattar-go/app/config"
	"github.com/kara9renai/yokattar-go/app/handler/httperror"
	"github.com/kara9renai/yokattar-go/app/handler/request"
)

// Handle Request for `GET /accounts/{username}/followers`
func (h *handler) Followers(w http.ResponseWriter, r *http.Request) {

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

	followers, err := a.FindFollowers(ctx, account.ID, limit)

	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	// maxId, err := request.URLParamOf(r, "max_id")

	// if err != nil {
	// 	maxId = config.DEFAULT_MAX_ID
	// }

	// sinceId, err := request.URLParamOf(r, "since_id")

	// if err != nil {
	// 	sinceId = config.DEFAULT_SINCE_ID
	// }

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(followers); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
