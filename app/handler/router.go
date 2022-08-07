package handler

import (
	"net/http"
	"time"

	"github.com/kara9renai/yokattar-go/app/app"
	"github.com/kara9renai/yokattar-go/app/handler/accounts"
	"github.com/kara9renai/yokattar-go/app/handler/health"

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

	r.Mount("/v1/accounts", accounts.NewRouter(app))
	r.Mount("/v1/health", health.NewRouter())

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
