package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/mcrgnt/yp1/internal/agent"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	agt, err := agent.NewAgentContext(ctx)
	if err != nil {
		log.Fatalf("new agent: %v", err)
	}
	agt.Run()
}
