package favorite

import (
	"encoding/json"
	"net/http"

	"github.com/kara9renai/yokattar-go/app/handler/auth"
	"github.com/kara9renai/yokattar-go/app/handler/httperror"
)

type FavoriteRequest struct {
	StatusId int `json:"favorite_id"`
}

// Handle Request for POST /v1/favorite
func (h *handler) Favorite(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	account := auth.AccountOf(r)
	var req FavoriteRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httperror.BadRequest(w, err)
		return
	}

	f := h.app.Dao.Favorite()
	isFavorite, err := f.Favorite(ctx, account.ID, int64(req.StatusId))
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}
	if !isFavorite {
		_, err := f.Confirm(ctx, account.ID, int64(req.StatusId))
		if err != nil {
			httperror.BadRequest(w, err)
			return
		}
	}
	favorite, err := f.Get(ctx, account.ID, int64(req.StatusId))
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(favorite); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
