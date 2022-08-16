package timelines

import (
	"encoding/json"
	"net/http"

	"github.com/kara9renai/yokattar-go/app/handler/httperror"
)

// Handle Request for `GET /timelines/public`
func (h *handler) GetPublic(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	t := h.app.Dao.Timeline()

	// 引数は固定
	statuses, err := t.GetPublicTimelines(ctx, 0, 0, 5)

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
