package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mcrgnt/yp1/internal/storage"

	"github.com/microgiantya/logger"
)

type API struct {
	srv     *http.Server
	ctx     *logger.Logger
	storage storage.MemStorage
}

type NewAPIParams struct {
	Ctx     context.Context
	Address string
	Storage storage.MemStorage
}

func NewAPI(params *NewAPIParams) (api *API) {
	api = &API{
		srv: &http.Server{
			Addr: params.Address,
		},
		ctx: logger.NewLoggerContext(params.Ctx, &logger.LoggerInitParams{
			Severity:       7,
			UniqueIDPrefix: "api",
			Version:        "v-",
		}),
		storage: params.Storage,
	}

	api.srv.Handler = NewDefaultHandler(&NewDefaultHandlerParams{
		Storage: params.Storage,
		Ctx:     params.Ctx,
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
	return t.srv.ListenAndServe()
}
