package statuses

import (
	"net/http"

	"github.com/kara9renai/yokattar-go/app/handler/httperror"
	"github.com/kara9renai/yokattar-go/app/handler/request"
)

// Handle Request for `DELETE /v1/statuses/:id`
func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ids, err := request.IDOf(r)

	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	s := h.app.Dao.Status() // domain/repositoryの取得

	err = s.DeleteStatus(ctx, ids)

	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}
