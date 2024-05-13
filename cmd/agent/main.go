package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/mcrgnt/yp1/internal/agent"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	agent, err := agent.NewAgent(ctx)
	if err != nil {
		panic(err)
	}
	agent.Run()
}
