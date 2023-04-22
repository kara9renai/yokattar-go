package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/kara9renai/yokattar-go/internal/app"
	"github.com/kara9renai/yokattar-go/pkg/config"
	"github.com/kara9renai/yokattar-go/pkg/server/handler"
	"golang.org/x/sync/errgroup"
)

type ApiServer struct {
	srv *http.Server
}

func (b *ApiServer) Init() {
	app, err := app.NewApp()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	addr := ":" + strconv.Itoa(config.Port())
	log.Printf("Serve on http://%s", addr)
	srv := &http.Server{
		Addr:    addr,
		Handler: handler.NewRouter(app),
	}
	b.srv = srv
}

func (b *ApiServer) Serve(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGTERM, os.Interrupt, syscall.SIGINT)
	defer stop()

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(
		func() error {
			if err := b.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				return err
			}
			return nil
		})

	<-ctx.Done()
	log.Println("Server Shutting down . . . ")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := b.srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("failed to shutdown: %+v", err)
	}
	log.Println("Server Shutdown")
	return eg.Wait()
}
