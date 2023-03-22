package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/kara9renai/yokattar-go/app/domain/object"
	"github.com/kara9renai/yokattar-go/app/server/handler/httperror"
)

type AddRequest struct {
	Username    string
	Password    string
	DisplayName *string
	Avatar      *string
	Header      *string
	Note        *string
}

// Handle Request for POST /v1/accounts
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httperror.BadRequest(w, err)
		return
	}
	account := new(object.Account)
	account.Username = req.Username
	account.DisplayName = req.DisplayName
	account.Avatar = req.Avatar
	account.Header = req.Header
	account.Note = req.Note

	if err := account.SetPassword(req.Password); err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	a := h.app.Dao.Account()
	account, err := a.Create(ctx, account)
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
