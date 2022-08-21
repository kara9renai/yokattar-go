package accounts

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kara9renai/yokattar-go/app/app"
	"github.com/kara9renai/yokattar-go/app/handler/auth"
)

type handler struct {
	app *app.App
}

func NewRouter(app *app.App) http.Handler {
	r := chi.NewRouter()
	h := &handler{app: app}
	r.Post("/", h.Create)
	r.Get("/{username}", h.Get)
	r.With(auth.Middleware(app)).Post("/{username}/follow", h.Follow)
	return r
}
