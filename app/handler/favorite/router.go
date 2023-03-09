package favorite

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
	r.With(auth.Middleware(app)).Post("/", h.Favorite)
	return r
}
