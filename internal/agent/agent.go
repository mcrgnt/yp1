package agent

import (
	"context"
	"fmt"
	"time"

	"github.com/mcrgnt/yp1/internal/agent/config"
	"github.com/mcrgnt/yp1/internal/metrics"
	"github.com/mcrgnt/yp1/internal/storage"
)

type Agent struct {
	Storage        storage.MemStorage
	address        string
	pollInterval   time.Duration
	reportInterval time.Duration
}

func NewAgentContext(ctx context.Context) (agent *Agent, err error) {
	agent = &Agent{}
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

func (t *Agent) Run(ctx context.Context) {
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

		case <-ctx.Done():
			return
		}
	}
}
