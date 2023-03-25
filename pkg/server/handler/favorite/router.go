package favorite

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
	r.With(middleware.Authenticate(app)).Post("/create", h.Create)
	return r
}
