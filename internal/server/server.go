package server

import (
	"context"
	"fmt"

	"github.com/mcrgnt/yp1/internal/api"
	"github.com/mcrgnt/yp1/internal/server/config"
	"github.com/mcrgnt/yp1/internal/storage"
)

type Server struct {
	api     *api.API
	address string
}

func NewServer() (server *Server, err error) {
	server = &Server{}
	cfg, err := config.NewConfig()
	if err != nil {
		return
	}

	server.address = cfg.Address
	server.api = api.NewAPI(&api.NewAPIParams{
		Address: cfg.Address,
		Storage: storage.NewStorage(&storage.NewMemStorageParams{
			Type: cfg.StorageType,
		}),
	})

	return
}

func (t *Server) Run(ctx context.Context) (chan struct{}, error) {
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
	close(graseful)
}
