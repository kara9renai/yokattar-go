package accounts

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kara9renai/yokattar-go/app/handler/httperror"
)

// Handle Request for `GET /v1/accounts/:username`
func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := chi.URLParam(r, "username")

	a := h.app.Dao.Account() // domain/repository の取得
	account, err := a.FindByUsername(ctx, username)

	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	if account == nil {
		err := errors.New("user not found")
		println(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
