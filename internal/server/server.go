package server

import (
	"context"
	"errors"
	"fmt"

	"github.com/mcrgnt/yp1/internal/agent/api"
	"github.com/mcrgnt/yp1/internal/filer"
	"github.com/mcrgnt/yp1/internal/server/internal/config"
	"github.com/mcrgnt/yp1/internal/store/store"
	"github.com/rs/zerolog"
)

type Server struct {
	cfg    *config.Config
	api    *api.API
	filer  *filer.Filer
	logger *zerolog.Logger
}

type NewServerParams struct {
	Logger *zerolog.Logger
}

func NewServerContext(ctx context.Context, params *NewServerParams) (*Server, error) {
	server := &Server{
		logger: params.Logger,
	}
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

func (t *Server) Run(ctx context.Context) (graceful chan struct{}, err error) {
	if t.cfg.Restore && t.cfg.FileStoragePath != "" {
		if e := t.filer.Read(); e != nil {
			err = fmt.Errorf("read failed: %w", e)
			return
		}
	}

	graceful = make(chan struct{})

	go func() {
		<-ctx.Done()
		if e := t.shutdown(context.Background(), graceful); e != nil {
			err = errors.Join(err, fmt.Errorf("srv shutdown failed: %w", e))
		}
	}()

	if err = t.api.Run(); err != nil {
		return graceful, fmt.Errorf("server run: %w", err)
	}
	return graceful, nil
}

func (t *Server) shutdown(ctx context.Context, graseful chan struct{}) (err error) {
	if e := t.api.Shutdown(ctx); e != nil {
		err = errors.Join(err, fmt.Errorf("srv shutdown failed: %w", e))
	}
	if t.cfg.FileStoragePath != "" {
		if e := t.filer.Write(); e != nil {
			err = errors.Join(err, fmt.Errorf("write failed: %w", e))
		}
	}
	close(graseful)
	return
}
