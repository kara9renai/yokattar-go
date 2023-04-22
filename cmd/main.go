package main

import (
	"context"
	"log"

	"github.com/kara9renai/yokattar-go/pkg/server"
)

func main() {
	srv := server.ApiServer{}
	srv.Init()
	if err := srv.Serve(context.Background()); err != nil {
		log.Fatalf("serve: %+v", err)
	}
}
