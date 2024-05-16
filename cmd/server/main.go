package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/mcrgnt/yp1/internal/server"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	srv, err := server.NewServer()
	if err != nil {
		log.Fatalf("new server: %v", err)
	}

	err = srv.Run(ctx)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
