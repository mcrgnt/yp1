package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/mcrgnt/yp1/internal/common"
	"github.com/mcrgnt/yp1/internal/compress/gzip"
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

func NewDefaultHandler(params *NewDefaultHandlerParams) (handler *DefaultHandler) {
	handler = &DefaultHandler{
		storage: params.Storage,
		R:       chi.NewRouter(),
		logger:  params.Logger,
	}

	handler.R.Group(func(r chi.Router) {
		// r.Use(handler.middlewareLogger)
		r.Use(middleware.Logger)
		r.Post("/update/{type}/{name}/{value}", handler.handlerUpdate)
	})
	handler.R.Group(func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Post("/update/{type}/", handler.handlerUpdate)
	})
	handler.R.Group(func(r chi.Router) {
		r.Use(middleware.Logger, middleware.Compress(common.CompressLevel, common.ContentTypeToCompressList...))
		r.Post("/update/", handler.handlerUpdateJSON)
	})
	handler.R.Group(func(r chi.Router) {
		r.Use(middleware.Logger, middleware.Compress(common.CompressLevel, common.ContentTypeToCompressList...))
		r.Post("/value/", handler.handlerValueJSON)
	})
	handler.R.Group(func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Get("/value/{type}/{name}", handler.handlerValue)
	})
	handler.R.Group(func(r chi.Router) {
		r.Use(middleware.Logger, middleware.Compress(common.CompressLevel, common.ContentTypeToCompressList...))
		r.Get("/", handler.handlerRoot)
	})
	return
}

func (t *DefaultHandler) CheckCompress(r *http.Request) (io.Reader, error) {
	if common.CheckContentEncodingGZIP(r) {
		if b, err := gzip.Decompress(r.Body); err != nil {
			return nil, fmt.Errorf("decompress failed: %w", err)
		} else {
			return b, nil
		}
	} else {
		return r.Body, nil
	}
}
