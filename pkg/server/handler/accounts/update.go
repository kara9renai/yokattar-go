package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/kara9renai/yokattar-go/pkg/domain/object"
	"github.com/kara9renai/yokattar-go/pkg/dto"
	"github.com/kara9renai/yokattar-go/pkg/http/middleware"
	"github.com/kara9renai/yokattar-go/pkg/server/handler/httperror"
)

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	diplayName := r.FormValue("display_name")
	avatar := r.FormValue("avatar")
	header := r.FormValue("header")
	note := r.FormValue("note")

	dto := dto.Credentials{
		DisplayName: diplayName,
		Avatar:      avatar,
		Header:      header,
		Note:        note,
	}

	targetUser := middleware.AccountOf(r)
	a := h.app.Dao.Account() // get domain/repository
	updateUser, err := a.Update(ctx, targetUser.ID, dto)
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
