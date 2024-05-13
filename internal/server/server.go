package server

import (
	"context"
	"sync"

	"github.com/mcrgnt/yp1/internal/api"
	"github.com/mcrgnt/yp1/internal/server/config"
	"github.com/mcrgnt/yp1/internal/storage"

	"github.com/microgiantya/logger"
)

const (
	logSeverity = 7
)

var (
	ServerVersion = "v-"
)

type Server struct {
	ctx     *logger.Logger
	api     *api.API
	wg      *sync.WaitGroup
	address string
}

func NewServerContext(ctx context.Context) (server *Server, err error) {
	server = &Server{
		ctx: logger.NewLoggerContext(ctx, &logger.LoggerInitParams{
			Severity:       logSeverity,
			UniqueIDPrefix: "srv",
			Version:        ServerVersion,
		}),
		wg: &sync.WaitGroup{},
	}

	cfg, err := config.NewConfig()
	if err != nil {
		return
	}

	server.address = cfg.Address
	server.api = api.NewAPI(ctx, &api.NewAPIParams{
		Address: cfg.Address,
		Storage: storage.NewMemStorage(&storage.NewMemStorageParams{
			Type: cfg.StorageType,
		}),
	})

	return
}

func (t *Server) report() {
	t.ctx.LogInformational("address: " + t.address)
}

func (t *Server) Run() {
	t.report()
	t.ctx.LogNotice("starting")
	t.wg.Add(1)
	go func() {
		t.wg.Done()
		<-t.ctx.Done()
		t.ctx.LogNotice("closing")
		t.api.Close()
	}()

	t.wg.Wait()
	_ = t.api.Run()
}
