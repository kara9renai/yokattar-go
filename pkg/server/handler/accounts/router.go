package accounts

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
	r.Post("/", h.Create)
	r.Get("/{username}", h.Get)
	r.With(middleware.Authenticate(app)).Post("/{username}/follow", h.Follow)
	r.Get("/{username}/following", h.Following)
	r.Get("/{username}/followers", h.Followers)
	r.With(middleware.Authenticate(app)).Post("/{username}/unfollow", h.Unfollow)
	r.With(middleware.Authenticate(app)).Post("/update_credentials", h.Update)
	return r
}
