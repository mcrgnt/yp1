package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/mcrgnt/yp1/internal/server"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	server, err := server.NewServer(ctx)
	if err != nil {
		panic(err)
	}
	server.Run()
}
