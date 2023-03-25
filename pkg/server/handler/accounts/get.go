package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/kara9renai/yokattar-go/pkg/server/handler/httperror"
	"github.com/kara9renai/yokattar-go/pkg/server/handler/request"
)

// Handle Request for `GET /v1/accounts/:username`
func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := request.UsernameOf(r)

	a := h.app.Dao.Account() // domain/repository の取得
	account, err := a.FindByUsername(ctx, username)

	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
