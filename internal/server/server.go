package server

import (
	"context"
	"flag"
	"fmt"
	"sync"

	"github.com/caarlos0/env/v11"
	"github.com/mcrgnt/yp1/internal/api"
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
	ctx         *logger.Logger
	api         *api.API
	wg          *sync.WaitGroup
	Address     string `env:"ADDRESS"`
	StorageType string `env:"MEMORY"`
}

func NewServer(ctx context.Context) (server *Server, err error) {
	server = &Server{
		ctx: logger.NewLoggerContext(ctx, &logger.LoggerInitParams{
			Severity:       logSeverity,
			UniqueIDPrefix: "srv",
			Version:        ServerVersion,
		}),
		wg: &sync.WaitGroup{},
	}

	err = server.paramsParse()
	if err != nil {
		return
	}

	server.api = api.NewAPI(ctx, &api.NewAPIParams{
		Address: server.Address,
		Storage: storage.NewMemStorage(&storage.NewMemStorageParams{
			Type: server.StorageType,
		}),
	})

	return
}

func (t *Server) paramsParseEnv() error {
	err := env.Parse(t)
	if err != nil {
		return fmt.Errorf("parse env: %v", err)
	}
	return nil
}

func (t *Server) paramsParseFlag() {
	flag.StringVar(&t.Address, "a", "localhost:8080", "")
	flag.Parse()
}

func (t *Server) paramsParse() error {
	t.paramsParseFlag()
	return t.paramsParseEnv()
}

func (t *Server) report() {
	t.ctx.LogInformational("address:" + t.Address)
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
