package timelines

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kara9renai/yokattar-go/internal/app"
	"github.com/kara9renai/yokattar-go/pkg/http/middleware"
)

type handler struct {
	app *app.App
}

func NewRouter(app *app.App) http.Handler {
	r := chi.NewRouter()

	h := &handler{app: app}

	r.Get("/public", h.Public)
	r.With(middleware.Authenticate(app)).Get("/home", h.Home)

	return r
}
