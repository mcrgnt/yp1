package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/mcrgnt/yp1/internal/agent"
	"github.com/rs/zerolog"
)

func main() {
	log := zerolog.New(zerolog.SyncWriter(os.Stdout)).Level(zerolog.DebugLevel)
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	if agt, err := agent.NewAgent(&agent.NewAgentParams{
		Logger: &log,
	}); err != nil {
		log.Fatal().Msgf("new agent failed: %v", err)
	} else {
		agt.Run(ctx)
	}
}
