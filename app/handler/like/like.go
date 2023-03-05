package like

import (
	"encoding/json"
	"net/http"

	"github.com/kara9renai/yokattar-go/app/handler/auth"
	"github.com/kara9renai/yokattar-go/app/handler/httperror"
)

type LikeRequest struct {
	StatusId int64
}

// Handle Request for POST /v1/like
func (h *handler) Like(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	account := auth.AccountOf(r)
	var req LikeRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httperror.BadRequest(w, err)
		return
	}

	l := h.app.Dao.Like()
	like, err := l.LikeByStatusId(ctx, account.ID, req.StatusId)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(like); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
