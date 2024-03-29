package handler

import (
	"net/http"
	"time"

	"github.com/kara9renai/yokattar-go/internal/app"
	mymiddleware "github.com/kara9renai/yokattar-go/pkg/http/middleware"
	"github.com/kara9renai/yokattar-go/pkg/server/handler/accounts"
	"github.com/kara9renai/yokattar-go/pkg/server/handler/favorites"
	"github.com/kara9renai/yokattar-go/pkg/server/handler/health"
	"github.com/kara9renai/yokattar-go/pkg/server/handler/media"
	"github.com/kara9renai/yokattar-go/pkg/server/handler/statuses"
	"github.com/kara9renai/yokattar-go/pkg/server/handler/timelines"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func NewRouter(app *app.App) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(newCORS().Handler)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(mymiddleware.AvailableIP())

	r.Mount("/v1/accounts", accounts.NewRouter(app))
	r.Mount("/v1/statuses", statuses.NewRouter(app))
	r.Mount("/v1/health", health.NewRouter())
	r.Mount("/v1/media", media.NewRouter(app))
	r.Mount("/v1/timelines", timelines.NewRouter(app))
	r.Mount("/v1/favorites", favorites.NewRouter(app))
	return r
}

func newCORS() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPut,
			http.MethodPatch,
			http.MethodPost,
			http.MethodDelete,
			http.MethodOptions,
		},
	})
}
