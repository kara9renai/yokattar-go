package favorites

import (
	"net/http"

	"github.com/kara9renai/yokattar-go/pkg/http/middleware"
	"github.com/kara9renai/yokattar-go/pkg/server/handler/httperror"
	"github.com/kara9renai/yokattar-go/pkg/server/handler/request"
)

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := request.IDOf(r)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}
	account := middleware.AccountOf(r)

	f := h.app.Dao.Favorite()
	if err := f.Delete(ctx, account.ID, id); err != nil {
		httperror.InternalServerError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
}
