package main

import (
	"context"
	"log"
)

func main() {
	b := Boot{}
	b.SetUp()
	if err := b.Serve(context.Background()); err != nil {
		log.Fatalf("serve: %+v", err)
	}
}
