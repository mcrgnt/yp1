package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mcrgnt/yp1/internal/store/models"
	"github.com/rs/zerolog"
)

type API struct {
	srv     *http.Server
	storage models.Storage
}

type NewAPIParams struct {
	Storage models.Storage
	Logger  *zerolog.Logger
	Address string
}

func NewAPI(params *NewAPIParams) (api *API) {
	api = &API{
		srv: &http.Server{
			Addr: params.Address,
		},
		storage: params.Storage,
	}

	api.srv.Handler = NewDefaultHandler(&NewDefaultHandlerParams{
		Storage: params.Storage,
		Logger:  params.Logger,
	}).R
	return
}

func (t *API) Shutdown(ctx context.Context) error {
	if err := t.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("srv shutdown failed: %w", err)
	}
	return nil
}

func (t *API) Run() error {
	err := t.srv.ListenAndServe()
	if err != nil {
		return fmt.Errorf("listen and serve: %w", err)
	}
	return nil
}
