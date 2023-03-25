package timelines

import (
	"encoding/json"
	"net/http"

	"github.com/kara9renai/yokattar-go/pkg/config"
	"github.com/kara9renai/yokattar-go/pkg/server/handler/httperror"
	"github.com/kara9renai/yokattar-go/pkg/server/handler/request"
)

// Handle Request for `GET /timelines/public`
func (h *handler) Public(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	limit, err := request.URLParamOf(r, "limit")

	if err != nil {
		limit = config.DEFAULT_LIMIT
	}

	if limit > config.MAX_LIMIT {
		limit = config.MAX_LIMIT
	}

	maxId, err := request.URLParamOf(r, "max_id")

	if err != nil {
		maxId = config.DEFAULT_MAX_ID
	}

	sinceId, err := request.URLParamOf(r, "since_id")

	if err != nil {
		sinceId = config.DEFAULT_SINCE_ID
	}

	t := h.app.Dao.Timeline()

	statuses, err := t.GetPublic(ctx, maxId, sinceId, limit)

	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(statuses); err != nil {
		httperror.InternalServerError(w, err)
		return
	}

}
