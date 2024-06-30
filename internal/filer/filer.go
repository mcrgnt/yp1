package filer

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/mcrgnt/yp1/internal/store/models"
	"github.com/rs/zerolog"
)

const (
	FilePermissions = 0o644
)

type Filer struct {
	storage  models.Storage
	logger   *zerolog.Logger
	filePath string
}

type NewFilerParams struct {
	Storage       models.Storage
	Logger        *zerolog.Logger
	FilePath      string
	WriteInterval time.Duration
}

func NewFilerContext(ctx context.Context, params *NewFilerParams) *Filer {
	filer := &Filer{
		logger:   params.Logger,
		filePath: params.FilePath,
		storage:  params.Storage,
	}
	go filer.periodicWriteContext(ctx, params.WriteInterval)
	return filer
}

func (t *Filer) periodicWriteContext(ctx context.Context, duration time.Duration) {
	if duration.Seconds() == 0 {
		emitter := t.storage.Emitter()
		for {
			select {
			case <-ctx.Done():
				return
			case <-emitter:
				if err := t.Write(); err != nil {
					t.logger.Error().Msgf("write failed: %s", err)
				}
			}
		}
	} else {
		tick := time.NewTicker(duration)
		for {
			select {
			case <-ctx.Done():
				tick.Stop()
				return
			case <-tick.C:
				if err := t.Write(); err != nil {
					t.logger.Error().Msgf("write failed: %s", err)
				}
			}
		}
	}
}

func (t *Filer) Read() error {
	if data, err := os.ReadFile(t.filePath); err != nil {
		return fmt.Errorf("read file failed: %w", err)
	} else {
		if err := t.storage.SetAllJSON(data); err != nil {
			return fmt.Errorf("set all json failed: %w", err)
		}
	}
	return nil
}

func (t *Filer) Write() error {
	if data, err := t.storage.GetAllJSON(); err != nil {
		return fmt.Errorf("get all json failed: %w", err)
	} else {
		if err := os.WriteFile(t.filePath, data, FilePermissions); err != nil {
			return fmt.Errorf("write file failed: %w", err)
		}
	}
	return nil
}
