package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/mcrgnt/yp1/internal/common"
	"github.com/mcrgnt/yp1/internal/storage"
	"github.com/rs/zerolog"
)

var (
	htmlHeader = `<!DOCTYPE html><html><head><title>Metrics</title></head><body>`
	htmlFooter = `</body></html>`
)

type DefaultHandler struct {
	storage storage.Storage
	R       *chi.Mux
	logger  *zerolog.Logger
}

type NewDefaultHandlerParams struct {
	Storage storage.Storage
	Logger  *zerolog.Logger
}

func (t *DefaultHandler) handlerUpdateJSON(w http.ResponseWriter, r *http.Request) {
	fmt.Println(*r)
	var (
		err           error
		statusHeader  = http.StatusOK
		storageParams = &storage.StorageParams{}
		returnBody    []byte
	)

	defer func() {
		if err != nil {
			switch {
			case errors.Is(err, common.ErrEmptyMetricName):
				statusHeader = http.StatusNotFound
			default:
				statusHeader = http.StatusBadRequest
			}
		}
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(statusHeader)
		_, _ = w.Write(returnBody)
	}()

	switch r.Header.Get("content-type") {
	case "application/json":
		if err = json.NewDecoder(r.Body).Decode(storageParams); err != nil {
			return
		}
		err = t.storage.GetMetricString(storageParams)
	default:
		err = errors.New("not found")
		return
	}

	if err = t.storage.MetricSet(storageParams); err != nil {
		return
	}
	if err = t.storage.GetMetricString(storageParams); err != nil {
		return
	}
	returnBody, err = json.Marshal(storageParams)
}

func (t *DefaultHandler) handlerUpdate(w http.ResponseWriter, r *http.Request) {
	var (
		err          error
		statusHeader = http.StatusOK
	)

	defer func() {
		if err != nil {
			switch {
			case errors.Is(err, common.ErrEmptyMetricName):
				statusHeader = http.StatusNotFound
			default:
				statusHeader = http.StatusBadRequest
			}
		}
		w.WriteHeader(statusHeader)
	}()
	storageParams := &storage.StorageParams{
		Type:  chi.URLParam(r, "type"),
		Name:  chi.URLParam(r, "name"),
		Value: chi.URLParam(r, "value"),
	}
	err = t.storage.MetricSet(storageParams)
}

func (t *DefaultHandler) handlerValueJSON(w http.ResponseWriter, r *http.Request) {
	var (
		err           error
		statusHeader  = http.StatusOK
		storageParams = &storage.StorageParams{}
		returnBody    []byte
	)

	defer func() {
		if err != nil {
			switch {
			case errors.Is(err, common.ErrMetricNotFound):
				statusHeader = http.StatusNotFound
			default:
				statusHeader = http.StatusBadRequest
			}
			return
		}
		w.WriteHeader(statusHeader)
		_, _ = w.Write(returnBody)
	}()

	switch r.Header.Get("content-type") {
	case "application/json":
		if err = json.NewDecoder(r.Body).Decode(storageParams); err != nil {
			return
		}
		err = t.storage.GetMetricString(storageParams)
	default:
		err = errors.New("not found")
		return
	}

	returnBody, err = json.Marshal(storageParams)
}

func (t *DefaultHandler) handlerValue(w http.ResponseWriter, r *http.Request) {
	var (
		err          error
		statusHeader = http.StatusOK
	)

	defer func() {
		w.WriteHeader(statusHeader)
	}()

	storageParams := &storage.StorageParams{
		Type: chi.URLParam(r, "type"),
		Name: chi.URLParam(r, "name"),
	}

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
	_, _ = fmt.Fprint(w, storageParams.String)
}

func (t *DefaultHandler) handlerRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(htmlHeader + t.storage.GetMetricAll() + htmlFooter))
}

func (t *DefaultHandler) midLogger(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		responseData := &responseData{
			status: 0,
			size:   0,
		}
		lw := &loggingResponseWriter{
			ResponseWriter: w,
			responseData:   responseData,
		}
		h.ServeHTTP(lw, r)

		t.logger.Info().Msgf("url (%s) method (%s) status (%d) duration (%v) size (%d)",
			r.RequestURI,
			r.Method,
			responseData.status,
			time.Since(start),
			responseData.size)
	})
}

func NewDefaultHandler(params *NewDefaultHandlerParams) (handler *DefaultHandler) {
	handler = &DefaultHandler{
		storage: params.Storage,
		R:       chi.NewRouter(),
		logger:  params.Logger,
	}

	handler.R.Post("/update/{type}/{name}/{value}", handler.midLogger(handler.handlerUpdate))
	handler.R.Post("/update/{type}/", handler.midLogger(handler.handlerUpdate))
	handler.R.Post("/update/", handler.midLogger(handler.handlerUpdateJSON))
	handler.R.Post("/value/", handler.midLogger(handler.handlerValueJSON))
	handler.R.Get("/value/{type}/{name}", handler.midLogger(handler.handlerValue))
	handler.R.Get("/", handler.midLogger(handler.handlerRoot))

	return
}
