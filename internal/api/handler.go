package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mcrgnt/yp1/internal/common"
	"github.com/mcrgnt/yp1/internal/storage"
)

var (
	htmlHeader = `<!DOCTYPE html><html><head><title>Metrics</title></head><body>`
	htmlFooter = `</body></html>`
)

type DefaultHandler struct {
	storage storage.Storage
	R       *chi.Mux
}

type NewDefaultHandlerParams struct {
	Storage storage.Storage
}

func (t *DefaultHandler) writeResponse(w http.ResponseWriter, _ *http.Request, statusHeader int, err error) {
	if err != nil {
		w.WriteHeader(statusHeader)
	}
}

func (t *DefaultHandler) handlerUpdate(w http.ResponseWriter, r *http.Request) {
	var (
		err          error
		statusHeader = http.StatusOK
	)
	defer func() {
		t.writeResponse(w, r, statusHeader, err)
	}()

	updateParams := &storage.StorageParams{
		Type:  chi.URLParam(r, "type"),
		Name:  chi.URLParam(r, "name"),
		Value: chi.URLParam(r, "value"),
	}

	err = t.storage.MetricSet(updateParams)
	if err != nil {
		switch {
		case errors.Is(err, common.ErrEmptyMetricName):
			statusHeader = http.StatusNotFound
		default:
			statusHeader = http.StatusBadRequest
		}
	}
}

func (t *DefaultHandler) handlerValue(w http.ResponseWriter, r *http.Request) {
	var (
		statusHeader = 200
		err          error
	)
	defer func() { t.writeResponse(w, r, statusHeader, err) }()

	storageParams := &storage.StorageParams{
		Type: chi.URLParam(r, "type"),
		Name: chi.URLParam(r, "name"),
	}

	err = t.storage.GetMetricStringByName(storageParams)
	if err != nil {
		switch {
		case errors.Is(err, common.ErrMetricNotFound):
			statusHeader = http.StatusNotFound
		default:
			statusHeader = http.StatusBadRequest
		}
	}
	_, _ = fmt.Fprint(w, storageParams.String)
}

func (t *DefaultHandler) handlerRoot(w http.ResponseWriter, r *http.Request) {
	var (
		statusHeader = 200
		err          error
	)
	defer t.writeResponse(w, r, statusHeader, err)
	_, _ = w.Write([]byte(htmlHeader + t.storage.GetMetricAll() + htmlFooter))
}

func NewDefaultHandler(params *NewDefaultHandlerParams) (handler *DefaultHandler) {
	handler = &DefaultHandler{
		storage: params.Storage,
		R:       chi.NewRouter(),
	}

	handler.R.Post("/update/{type}/{name}/{value}", handler.handlerUpdate)
	handler.R.Get("/value/{type}/{name}", handler.handlerValue)
	handler.R.Get("/", handler.handlerRoot)

	return
}
