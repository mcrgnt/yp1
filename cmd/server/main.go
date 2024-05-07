package main

import (
	"context"
	"os"
	"os/signal"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	server, err := NewServer(ctx)
	if err != nil {
		panic(err)
	}
	server.Run()
}
