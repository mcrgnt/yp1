package server

import (
	"context"
	"fmt"

	"github.com/mcrgnt/yp1/internal/agent/api"
	"github.com/mcrgnt/yp1/internal/filer"
	"github.com/mcrgnt/yp1/internal/server/internal/config"
	"github.com/mcrgnt/yp1/internal/store/store"
	"github.com/rs/zerolog"
)

type Server struct {
	cfg   *config.Config
	api   *api.API
	filer *filer.Filer
}

type NewServerParams struct {
	Logger *zerolog.Logger
}

func NewServerContext(ctx context.Context, params *NewServerParams) (*Server, error) {
	server := &Server{}
	if cfg, err := config.NewConfig(); err != nil {
		return nil, fmt.Errorf("new server context failed: %w", err)
	} else {
		server.cfg = cfg
		strg := store.NewStorage(&store.NewStorageParams{
			Type: server.cfg.StorageType,
		})
		server.api = api.NewAPI(&api.NewAPIParams{
			Address: server.cfg.Address,
			Storage: strg,
			Logger:  params.Logger,
		})
		server.filer = filer.NewFilerContext(ctx, &filer.NewFilerParams{
			FilePath:      server.cfg.FileStoragePath,
			Logger:        params.Logger,
			WriteInterval: server.cfg.Parsed.StoreInterval,
			Storage:       strg,
		})
	}
	return server, nil
}

func (t *Server) Run(ctx context.Context) (chan struct{}, error) {
	if t.cfg.Restore && t.cfg.FileStoragePath != "" {
		t.filer.Read()
	}

	graseful := make(chan struct{})

	go func() {
		<-ctx.Done()
		t.shutdown(context.Background(), graseful)
	}()

	err := t.api.Run()
	if err != nil {
		return graseful, fmt.Errorf("server run: %w", err)
	}
	return graseful, nil
}

func (t *Server) shutdown(ctx context.Context, graseful chan struct{}) {
	t.api.Shutdown(ctx)
	if t.cfg.FileStoragePath != "" {
		t.filer.Write()
	}
	close(graseful)
}
