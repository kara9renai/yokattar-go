package timelines

import (
	"encoding/json"
	"net/http"

	"github.com/kara9renai/yokattar-go/app/config"
	"github.com/kara9renai/yokattar-go/app/handler/auth"
	"github.com/kara9renai/yokattar-go/app/handler/httperror"
	"github.com/kara9renai/yokattar-go/app/handler/request"
)

// Handle Request for `GET /timelines/home`
func (h *handler) Home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	account := auth.AccountOf(r)

	limit, err := request.URLParamOf(r, "limit")

	if err != nil {
		limit = config.DEFAULT_LIMIT
	}

	if limit > config.MAX_LIMIT {
		limit = config.MAX_LIMIT
	}

	t := h.app.Dao.Timeline() // domain/repository の取得

	statuses, err := t.GetHome(ctx, account.ID, limit)

	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(statuses); err != nil {
		httperror.InternalServerError(w, err)
		return
	}

}
