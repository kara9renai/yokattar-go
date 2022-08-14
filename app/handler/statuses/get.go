package statuses

import (
	"encoding/json"
	"net/http"

	"github.com/kara9renai/yokattar-go/app/handler/httperror"
	"github.com/kara9renai/yokattar-go/app/handler/request"
)

// Handle Request for `GET /v1/statuses/:id`
func (h *handler) Get(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	// IDの取得
	id, err := request.IDOf(r)

	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	s := h.app.Dao.Status() // domain/repositoryの取得

	status, err := s.GetStatus(ctx, id)

	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	a := h.app.Dao.Account() // domain/repositoryの取得

	account, err := a.FindByID(ctx, status.AccountID)

	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	status.Account = *account

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(status); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
