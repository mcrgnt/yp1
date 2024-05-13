package agent

import (
	"context"
	"fmt"
	"time"

	"github.com/mcrgnt/yp1/internal/agent/config"
	"github.com/mcrgnt/yp1/internal/metrics"
	"github.com/mcrgnt/yp1/internal/storage"
	"github.com/microgiantya/logger"
)

const (
	logSeverity = 7
)

var (
	AgentVersion = "v-"
)

type Agent struct {
	Storage        storage.MemStorage
	ctx            *logger.Logger
	address        string
	pollInterval   time.Duration
	reportInterval time.Duration
}

func NewAgentContext(ctx context.Context) (agent *Agent, err error) {
	agent = &Agent{
		ctx: logger.NewLoggerContext(ctx, &logger.LoggerInitParams{
			Severity:       logSeverity,
			UniqueIDPrefix: "srv",
			Version:        AgentVersion,
		}),
	}

	cfg, err := config.NewConfig()
	if err != nil {
		return
	}

	agent.address = cfg.Address
	agent.Storage = storage.NewMemStorage(&storage.NewMemStorageParams{
		Type: cfg.StorageType,
	})
	agent.pollInterval, err = time.ParseDuration(cfg.PollInterval + "s")
	if err != nil {
		err = fmt.Errorf("parse pollInterval: %w", err)
		return
	}
	agent.reportInterval, err = time.ParseDuration(cfg.ReportInterval + "s")
	if err != nil {
		err = fmt.Errorf("parse reportInterval: %w", err)
	}

	return
}

func (t *Agent) report() {
	t.ctx.LogInformational(fmt.Sprintf("address: %v", t.address))
	t.ctx.LogInformational(fmt.Sprintf("poll interval: %v", t.pollInterval))
	t.ctx.LogInformational(fmt.Sprintf("report interval: %v", t.reportInterval))
}

func (t *Agent) Run() {
	t.report()
	t.ctx.LogNotice("starting")
	metrics.PollMetrics(&metrics.PollMetricsParams{
		Storage: t.Storage,
	})
	metrics.ReportMetrics(&metrics.ReportMetricsParams{
		Storage: t.Storage,
		Address: t.address,
	})

	pollTicker := time.NewTicker(t.pollInterval)
	reportTicker := time.NewTicker(t.reportInterval)

	for {
		select {
		case <-pollTicker.C:
			metrics.PollMetrics(&metrics.PollMetricsParams{
				Storage: t.Storage,
			})
		case <-reportTicker.C:
			metrics.ReportMetrics(&metrics.ReportMetricsParams{
				Storage: t.Storage,
				Address: t.address,
			})

		case <-t.ctx.Done():
			t.ctx.LogNotice("stopping")
			return
		}
	}
}
