package api

import (
	"fmt"
	"net/http"

	"github.com/mcrgnt/yp1/internal/storage"
)

type API struct {
	srv     *http.Server
	storage storage.MemStorage
}

type NewAPIParams struct {
	Storage storage.MemStorage
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
	}).R
	return
}

func (t *API) Close() {
	_ = t.srv.Close()
}

func (t *API) Run() error {
	err := t.srv.ListenAndServe()
	if err != nil {
		return fmt.Errorf("listen and server: %w", err)
	}
	return nil
}
