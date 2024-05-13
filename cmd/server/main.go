package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/mcrgnt/yp1/internal/server"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	srv, err := server.NewServerContext(ctx)
	if err != nil {
		log.Fatalf("new server: %v", err)
	}
	srv.Run()
}
