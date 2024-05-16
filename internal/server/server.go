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
		Storage: storage.NewMemStorage(&storage.NewMemStorageParams{
			Type: cfg.StorageType,
		}),
	})

	return
}

func (t *Server) Run(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		t.api.Close()
	}()

	err := t.api.Run()
	if err != nil {
		return fmt.Errorf("server run: %w", err)
	}

	return nil
}
