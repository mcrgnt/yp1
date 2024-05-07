package main

import (
	"context"
	"os"
	"os/signal"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	agent, err := NewAgent(ctx)
	if err != nil {
		panic(err)
	}
	agent.Run()
}
