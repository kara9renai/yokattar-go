package main

import (
	"context"
	"log"

	"github.com/kara9renai/yokattar-go/pkg/server"
)

func main() {
	b := server.Boot{}
	b.SetUp()
	if err := b.Serve(context.Background()); err != nil {
		log.Fatalf("serve: %+v", err)
	}
}
