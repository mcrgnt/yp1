package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mcrgnt/yp1/internal/storage"

	"github.com/microgiantya/logger"
)

const (
	logSeverity = 7
)

type API struct {
	srv     *http.Server
	ctx     *logger.Logger
	storage storage.MemStorage
}

type NewAPIParams struct {
	Storage storage.MemStorage
	Address string
}

func NewAPI(ctx context.Context, params *NewAPIParams) (api *API) {
	api = &API{
		srv: &http.Server{
			Addr: params.Address,
		},
		ctx: logger.NewLoggerContext(ctx, &logger.LoggerInitParams{
			Severity:       logSeverity,
			UniqueIDPrefix: "api",
			Version:        "v-",
		}),
		storage: params.Storage,
	}

	api.srv.Handler = NewDefaultHandler(ctx, &NewDefaultHandlerParams{
		Storage: params.Storage,
	}).R
	return
}

func (t *API) Close() {
	err := t.srv.Close()
	if err != nil {
		t.ctx.LogError(fmt.Sprintf("api close: %v", err))
	}
}

func (t *API) Run() error {
	return fmt.Errorf("api run: %w", t.srv.ListenAndServe())
}
