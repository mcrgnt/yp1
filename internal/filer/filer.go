package filer

import (
	"context"
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
				t.Write()
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
				t.Write()
			}
		}
	}
}

func (t *Filer) Read() {
	if data, err := os.ReadFile(t.filePath); err != nil {
		t.logger.Error().Msgf("read file failed: %s", err)
	} else {
		if err := t.storage.SetAllJSON(data); err != nil {
			t.logger.Error().Msgf("set all json failed: %s", err)
		}
	}
}

func (t *Filer) Write() {
	if data, err := t.storage.GetAllJSON(); err != nil {
		t.logger.Error().Msgf("get all json failed: %s", err)
	} else {
		if err := os.WriteFile(t.filePath, data, FilePermissions); err != nil {
			t.logger.Error().Msgf("write file failed: %s", err)
		}
	}
}
