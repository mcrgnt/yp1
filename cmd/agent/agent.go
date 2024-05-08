package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/mcrgnt/yp1/internal/metrics"
	"github.com/mcrgnt/yp1/internal/storage"
	"github.com/microgiantya/logger"
)

var (
	AgentVersion = "v-"
)

type Agent struct {
	Address     string // `env:"ADDRESS"`
	StorageType string `env:"memory"`
	ctx         *logger.Logger
	//
	PollInterval   string // `env:"POLL_INTERVAL"`
	pollInterval   time.Duration
	ReportInterval string // `env:"REPORT_INTERVAL"`
	reportInterval time.Duration
	Storage        storage.MemStorage
}

func NewAgent(ctx context.Context) (agent *Agent, err error) {
	agent = &Agent{
		ctx: logger.NewLoggerContext(ctx, &logger.LoggerInitParams{
			Severity:       7,
			UniqueIDPrefix: "srv",
			Version:        AgentVersion,
		}),
		Storage: storage.NewMemStorage(&storage.NewMemStorageParams{}),
	}
	agent.paramsParseFlag()
	// err = agent.paramsParseEnv()
	// if err != nil {
	// 	return
	// }
	agent.pollInterval, err = time.ParseDuration(agent.PollInterval + "s")
	if err != nil {
		err = fmt.Errorf("parse pollInterval: %v", err)
		return
	}

	agent.reportInterval, err = time.ParseDuration(agent.ReportInterval + "s")
	if err != nil {
		err = fmt.Errorf("parse reportInterval: %v", err)
	}

	return
}

// func (t *Agent) paramsParseEnv() error {
// 	return env.Parse(t)
// }

func (t *Agent) paramsParseFlag() {
	flag.StringVar(&t.Address, "a", "localhost:8080", "")
	flag.StringVar(&t.PollInterval, "p", "2", "")
	flag.StringVar(&t.ReportInterval, "r", "10", "")
	flag.Parse()
}

func (t *Agent) report() {
	t.ctx.LogInformational(fmt.Sprintf("address: %v", t.Address))
	t.ctx.LogInformational(fmt.Sprintf("poll interval: %v", t.PollInterval))
	t.ctx.LogInformational(fmt.Sprintf("report interval: %v", t.ReportInterval))
}

func (t *Agent) Run() (err error) {
	t.report()
	t.ctx.LogNotice("starting")
	metrics.PollMetrics(&metrics.PollMetricsParams{
		Storage: t.Storage,
	})
	go metrics.ReportMetrics(&metrics.ReportMetricsParams{
		Storage: t.Storage,
		Address: t.Address,
	})

	pollTicker := time.NewTicker(t.pollInterval)
	reportTicker := time.NewTicker(t.reportInterval)

	for {
		select {
		case <-pollTicker.C:
			go metrics.PollMetrics(&metrics.PollMetricsParams{
				Storage: t.Storage,
			})
		case <-reportTicker.C:
			go metrics.ReportMetrics(&metrics.ReportMetricsParams{
				Storage: t.Storage,
				Address: t.Address,
			})

		case <-t.ctx.Done():
			t.ctx.LogNotice("stopping")
			return
		}
	}
}
