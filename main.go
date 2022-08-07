package main

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/kara9renai/yokattar-go/app/app"
	"github.com/kara9renai/yokattar-go/app/config"
	"github.com/kara9renai/yokattar-go/app/handler"
)

func main() {
	log.Fatalf("%+v", serve(context.Background()))
}

func serve(ctx context.Context) error {
	app, err := app.NewApp()
	if err != nil {
		return err
	}
	addr := ":" + strconv.Itoa(config.Port())
	log.Printf("Serve on http://%s", addr)

	return http.ListenAndServe(addr, handler.NewRouter(app))
}
