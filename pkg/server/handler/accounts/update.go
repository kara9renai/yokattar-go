package accounts

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kara9renai/yokattar-go/pkg/http/middleware"
	"github.com/kara9renai/yokattar-go/pkg/server/handler/httperror"
)

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	displayName := r.FormValue("display_name")
	avatar := r.FormValue("avatar")
	header := r.FormValue("header")
	note := r.FormValue("note")
	log.Println("test:", avatar, header)

	username := middleware.AccountOf(r)
	username.DisplayName = &displayName
	username.Note = &note

	a := h.app.Dao.Account()
	updateUser, err := a.Update(ctx, username)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(updateUser); err != nil {
		httperror.InternalServerError(w, err)
	}
}
