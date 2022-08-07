package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/kara9renai/yokattar-go/app/domain/object"
	"github.com/kara9renai/yokattar-go/app/handler/httperror"
)

type AddRequest struct {
	Username string
	Password string
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {

	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httperror.BadRequest(w, err)
		return
	}
	account := new(object.Account)
	account.Username = req.Username
	if err := account.SetPassword(req.Password); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
	_ = h.app.Dao.Account()
	panic("Must Implement Account Registration")

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
