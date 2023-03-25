package statuses

import (
	"encoding/json"
	"net/http"

	"github.com/kara9renai/yokattar-go/pkg/http/middleware"
	"github.com/kara9renai/yokattar-go/pkg/server/handler/httperror"
)

type AddRequest struct {
	Status string
}

// Handle Request for `POST /v1/statuses`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	account := middleware.AccountOf(r)

	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httperror.BadRequest(w, err)
		return
	}

	s := h.app.Dao.Status()
	status, err := s.Create(ctx, account.ID, req.Status)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}
	if status != nil {
		a := h.app.Dao.Account()
		account, err := a.FindByID(ctx, status.AccountID)
		if err != nil {
			httperror.InternalServerError(w, err)
			return
		}

		status.Account = *account
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
