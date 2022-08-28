package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kara9renai/yokattar-go/app/handler/httperror"
)

// Handle Request for `GET /accounts/{username}/following`
func (h *handler) Following(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	username := chi.URLParam(r, "username")

	a := h.app.Dao.Account() // domain/repository の取得

	account, err := a.FindByUsername(ctx, username)

	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	followingUsers, err := a.FindFollowing(ctx, account.ID)

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
