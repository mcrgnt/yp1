package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/mcrgnt/yp1/internal/server"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	srv, err := server.NewServerContext(ctx)
	if err != nil {
		panic(err)
	}
	srv.Run()
}
