package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mcrgnt/yp1/internal/storage"
)

var (
	htmlHeader = `<!DOCTYPE html><html><head><title>Metrics</title></head><body>`
	htmlFooter = `</body></html>`
)

type DefaultHandler struct {
	storage storage.MemStorage
	R       *chi.Mux
}

type NewDefaultHandlerParams struct {
	Storage storage.MemStorage
}

func (t *DefaultHandler) writeResponse(w http.ResponseWriter, _ *http.Request, statusHeader int, err error) {
	if err != nil {
		w.WriteHeader(statusHeader)
	}
}

func (t *DefaultHandler) handlerUpdate(w http.ResponseWriter, r *http.Request) {
	var (
		statusHeader = 200
		err          error
	)
	defer func() { t.writeResponse(w, r, statusHeader, err) }()

	updateParams := &storage.StorageParams{
		Type:  chi.URLParam(r, "type"),
		Name:  chi.URLParam(r, "name"),
		Value: chi.URLParam(r, "value"),
	}

	err = updateParams.ValidateType()
	if err != nil {
		statusHeader = http.StatusBadRequest
		return
	}

	err = updateParams.ValidateName()
	if err != nil {
		statusHeader = http.StatusNotFound
		return
	}

	err = updateParams.ValidateValue()
	if err != nil {
		statusHeader = http.StatusBadRequest
		return
	}

	statusHeader = http.StatusOK
	t.storage.Update(updateParams)
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

	err = storageParams.ValidateType()
	if err != nil {
		statusHeader = http.StatusBadRequest
		return
	}
	err = storageParams.ValidateName()
	if err != nil {
		statusHeader = http.StatusNotFound
		return
	}

	value, err := t.storage.GetByType(storageParams)
	if err != nil {
		statusHeader = http.StatusNotFound
		return
	}
	_, _ = fmt.Fprintf(w, "%s", value)
}

func (t *DefaultHandler) handlerRoot(w http.ResponseWriter, r *http.Request) {
	var (
		statusHeader = 200
		err          error
	)
	defer t.writeResponse(w, r, statusHeader, err)
	_, _ = w.Write([]byte(htmlHeader + t.storage.GetAll() + htmlFooter))
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
