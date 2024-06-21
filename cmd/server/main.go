package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"

	"github.com/mcrgnt/yp1/internal/server"
	"github.com/rs/zerolog"
)

func main() {
	log := zerolog.New(zerolog.SyncWriter(os.Stdout)).Level(zerolog.DebugLevel)

	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	srv, err := server.NewServerContext(ctx, &server.NewServerParams{
		Logger: &log,
	})
	if err != nil {
		log.Fatal().Msgf("new server: %v", err)
	}

	graseful, err := srv.Run(ctx)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal().Msgf("run server: %v", err)
	}
	<-graseful
}
