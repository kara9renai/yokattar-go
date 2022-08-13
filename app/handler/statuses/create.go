package statuses

import (
	"encoding/json"
	"net/http"

	"github.com/kara9renai/yokattar-go/app/handler/auth"
	"github.com/kara9renai/yokattar-go/app/handler/httperror"
)

type AddRequest struct {
	Status string
}

// Handle Request for `POST /v1/statuses`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	account := auth.AccountOf(r)

	var req AddRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httperror.BadRequest(w, err)
		return
	}

	s := h.app.Dao.Status()
	status, err := s.CreateStatus(ctx, account.ID, req.Status)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
