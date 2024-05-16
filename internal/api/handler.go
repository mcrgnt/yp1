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

func (t *DefaultHandler) handlerUpdate(w http.ResponseWriter, r *http.Request) {
	var (
		err          error
		statusHeader = http.StatusOK
	)
	defer func() {
		if err != nil {
			w.WriteHeader(statusHeader)
		}
	}()

	storageParams := &storage.StorageParams{
		Type:  chi.URLParam(r, "type"),
		Name:  chi.URLParam(r, "name"),
		Value: chi.URLParam(r, "value"),
	}

	err = t.storage.MetricSet(storageParams)
	if err != nil {
		switch {
		case errors.Is(err, common.ErrEmptyMetricName):
			statusHeader = http.StatusNotFound
		default:
			statusHeader = http.StatusBadRequest
		}
	}
	fmt.Printf("----- update: %s %v %+v %v %d <<\n", r.Method, r.URL.Path, *storageParams, err, statusHeader)
}

func (t *DefaultHandler) handlerValue(w http.ResponseWriter, r *http.Request) {
	var (
		statusHeader = http.StatusOK
		err          error
	)
	defer func() {
		if err != nil {
			w.WriteHeader(statusHeader)
		}
	}()

	storageParams := &storage.StorageParams{
		Type: chi.URLParam(r, "type"),
		Name: chi.URLParam(r, "name"),
	}

	defer fmt.Printf("----- value: %s %v %+v %v %d<<\n", r.Method, r.URL.Path, *storageParams, err, statusHeader)

	err = t.storage.GetMetricString(storageParams)
	if err != nil {
		switch {
		case errors.Is(err, common.ErrMetricNotFound):
			statusHeader = http.StatusNotFound
		default:
			statusHeader = http.StatusBadRequest
		}
		return
	}
	fmt.Println(storageParams.String)
	_, _ = fmt.Fprint(w, storageParams.String)
}

func (t *DefaultHandler) handlerRoot(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(htmlHeader + t.storage.GetMetricAll() + htmlFooter))
}

func NewDefaultHandler(params *NewDefaultHandlerParams) (handler *DefaultHandler) {
	handler = &DefaultHandler{
		storage: params.Storage,
		R:       chi.NewRouter(),
	}

	handler.R.Post("/update/{type}/{name}/{value}", handler.handlerUpdate)
	handler.R.Post("/update/{type}/", handler.handlerUpdate)
	handler.R.Get("/value/{type}/{name}", handler.handlerValue)
	handler.R.Get("/", handler.handlerRoot)

	return
}
