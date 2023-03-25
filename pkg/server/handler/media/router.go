package media

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kara9renai/yokattar-go/internal/app"
)

type handler struct {
	app *app.App
}

func NewRouter(app *app.App) http.Handler {
	r := chi.NewRouter()
	h := &handler{app: app}
	r.Post("/", h.Upload)
	return r
}
