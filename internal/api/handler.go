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

func (t *DefaultHandler) writeResponse(w http.ResponseWriter, statusHeader int, err error) {
	if err != nil {
		fmt.Println(">>>>>", err)
		w.WriteHeader(statusHeader)
	}
}

func (t *DefaultHandler) handlerUpdate(w http.ResponseWriter, r *http.Request) {
	var (
		err          error
		statusHeader = http.StatusOK
	)
	defer func() {
		t.writeResponse(w, statusHeader, err)
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
			fmt.Printf("update: %+v %v %d <<\n", *updateParams, err, statusHeader)
		default:
			statusHeader = http.StatusBadRequest
			fmt.Printf("update: %+v %v %d <<\n", *updateParams, err, statusHeader)
		}
	}
}

func (t *DefaultHandler) handlerValue(w http.ResponseWriter, r *http.Request) {
	var (
		statusHeader = 200
		err          error
	)
	defer func() {
		t.writeResponse(w, statusHeader, err)
	}()

	storageParams := &storage.StorageParams{
		Type: chi.URLParam(r, "type"),
		Name: chi.URLParam(r, "name"),
	}

	err = t.storage.GetMetricStringByName(storageParams)
	if err != nil {
		fmt.Println(err)
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
	defer t.writeResponse(w, statusHeader, err)
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
