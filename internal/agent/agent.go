package agent

import (
	"context"
	"fmt"
	"time"

	"github.com/mcrgnt/yp1/internal/agent/internal/config"
	"github.com/mcrgnt/yp1/internal/metrics"
	"github.com/mcrgnt/yp1/internal/store/models"
	"github.com/mcrgnt/yp1/internal/store/store"
	"github.com/rs/zerolog"
)

type Agent struct {
	Storage        models.Storage
	logger         *zerolog.Logger
	address        string
	pollInterval   time.Duration
	reportInterval time.Duration
}

type NewAgentParams struct {
	Logger *zerolog.Logger
}

func NewAgent(params *NewAgentParams) (*Agent, error) {
	agent := &Agent{
		logger: params.Logger,
	}

	if cfg, err := config.NewConfig(); err != nil {
		return nil, fmt.Errorf("new config failed: %w", err)
	} else {
		agent.address = cfg.Address
		agent.Storage = store.NewStorage(&store.NewStorageParams{
			Type: cfg.StorageType,
		})
		if agent.pollInterval, err = time.ParseDuration(cfg.PollInterval + "s"); err != nil {
			return nil, fmt.Errorf("parse pollInterval failed: %w", err)
		}
		if agent.reportInterval, err = time.ParseDuration(cfg.ReportInterval + "s"); err != nil {
			return nil, fmt.Errorf("parse reportInterval failed: %w", err)
		}
	}
	return agent, nil
}

func (t *Agent) Run(ctx context.Context) error {
	pollTicker := time.NewTicker(t.pollInterval)
	reportTicker := time.NewTicker(t.reportInterval)

	if err := metrics.PollMetrics(&metrics.PollMetricsParams{
		Storage: t.Storage,
	}); err != nil {
		return fmt.Errorf("poll metrics failed: %w", err)
	}

	for {
		select {
		case <-pollTicker.C:
			if err := metrics.PollMetrics(&metrics.PollMetricsParams{
				Storage: t.Storage,
			}); err != nil {
				return fmt.Errorf("poll metrics failed: %w", err)
			}
		case <-reportTicker.C:
			if err := metrics.ReportMetrics(&metrics.ReportMetricsParams{
				Storage: t.Storage,
				Address: t.address,
			}); err != nil {
				return fmt.Errorf("report metrics failed: %w", err)
			}
		case <-ctx.Done():
			return nil
		}
	}
}
