package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/kara9renai/yokattar-go/app/app"
	"github.com/kara9renai/yokattar-go/app/config"
	"github.com/kara9renai/yokattar-go/app/handler"
	"golang.org/x/sync/errgroup"
)

func main() {
	if err := serve(context.Background()); err != nil {
		log.Fatalf("serve: %+v", err)
	}
}

func serve(ctx context.Context) error {
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

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGTERM, os.Interrupt, syscall.SIGINT)
	defer stop()

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(
		func() error {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				return err
			}
			return nil
		})

	<-ctx.Done()
	log.Println("Server Shutting down . . . ")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("failed to shutdown: %+v", err)
	}
	log.Println("Server Shutdown")
	return eg.Wait()
}
