package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/kara9renai/yokattar-go/pkg/domain/object"
	"github.com/kara9renai/yokattar-go/pkg/http/middleware"
	"github.com/kara9renai/yokattar-go/pkg/server/handler/httperror"
)

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	targetUser := middleware.AccountOf(r)
	diplayName := r.FormValue("display_name")
	if len(diplayName) != 0 {
		targetUser.DisplayName = &diplayName
	}
	avatar := r.FormValue("avatar")
	if len(avatar) != 0 {
		targetUser.Avatar = &avatar
	}
	header := r.FormValue("header")
	if len(header) != 0 {
		targetUser.Header = &header
	}
	note := r.FormValue("note")
	if len(note) != 0 {
		targetUser.Note = &note
	}

	a := h.app.Dao.Account() // get domain/repository
	updateUser, err := a.Update(ctx, targetUser)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(updateUser); err != nil {
		httperror.InternalServerError(w, err)
	}
}

func Test(r *http.Request, val string, e *object.Account) {
	newValue := r.FormValue(val)
	if len(newValue) != 0 {
		e.DisplayName = &newValue
	}
}
