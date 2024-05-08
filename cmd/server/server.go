package main

import (
	"context"
	"flag"
	"fmt"
	"sync"

	"github.com/caarlos0/env/v10"
	"github.com/mcrgnt/yp1/internal/api"
	"github.com/mcrgnt/yp1/internal/storage"

	"github.com/microgiantya/logger"
)

var (
	ServerVersion = "v-"
)

type Server struct {
	Address     string // `env:"ADDRESS"`
	StorageType string `env:"memory"`
	ctx         *logger.Logger
	api         *api.API
	wg          *sync.WaitGroup
}

func NewServer(ctx context.Context) (server *Server, err error) {
	server = &Server{
		ctx: logger.NewLoggerContext(ctx, &logger.LoggerInitParams{
			Severity:       7,
			UniqueIDPrefix: "srv",
			Version:        ServerVersion,
		}),
		wg: &sync.WaitGroup{},
	}
	server.paramsParseFlag()
	err = server.paramsParseEnv()
	if err != nil {
		return
	}

	server.api = api.NewAPI(&api.NewAPIParams{
		Ctx:     ctx,
		Address: server.Address,
		Storage: storage.NewMemStorage(&storage.NewMemStorageParams{
			Type: server.StorageType,
		}),
	})

	return
}

func (t *Server) paramsParseEnv() error {
	return env.Parse(t)
}

func (t *Server) paramsParseFlag() {
	flag.StringVar(&t.Address, "a", "localhost:8080", "")
	flag.Parse()
}

func (t *Server) report() {
	t.ctx.LogInformational(fmt.Sprintf("address: %s", t.Address))
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
