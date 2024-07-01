package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/mcrgnt/yp1/internal/compress/gzip"
	"github.com/mcrgnt/yp1/internal/store/models"
	"github.com/rs/zerolog"
)

const (
	contentType     = "Content-Type"
	contentEncoding = "Content-Encoding"
	acceptEncoding  = "Accept-Encoding"

	applicationJSON = "application/json"
	textHTML        = "text/html"
	gZip            = "gzip"

	compressLevel = 5
)

var (
	htmlHeader = `<!DOCTYPE html><html><head><title>Metrics</title></head><body>`
	htmlFooter = `</body></html>`

	contentTypeToCompressList = []string{"text/html", "application/json"}
)

func checkContentEncodingGZIP(r *http.Request) bool {
	return strings.Contains(r.Header.Get(contentEncoding), strings.ToLower(gZip))
}

func checkAcceptEncodingGZIP(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.Header.Get(contentEncoding), strings.ToLower(gZip)) {
		w.Header().Set(contentEncoding, gZip)
	}
}

type DefaultHandler struct {
	storage models.Storage
	R       *chi.Mux
	logger  *zerolog.Logger
}

type NewDefaultHandlerParams struct {
	Storage models.Storage
	Logger  *zerolog.Logger
}

func NewDefaultHandler(params *NewDefaultHandlerParams) (handler *DefaultHandler) {
	handler = &DefaultHandler{
		storage: params.Storage,
		R:       chi.NewRouter(),
		logger:  params.Logger,
	}
	handler.R.Use(middleware.Logger)

	handler.R.Group(func(r chi.Router) {
		r.Post("/update/{type}/{name}/{value}", handler.handlerUpdate)
		r.Post("/update/{type}/", handler.handlerUpdate)
		r.Get("/value/{type}/{name}", handler.handlerValue)
	})
	handler.R.Group(func(r chi.Router) {
		r.Use(middleware.Compress(compressLevel, contentTypeToCompressList...))
		r.Post("/update/", handler.handlerUpdateJSON)
		r.Post("/value/", handler.handlerValueJSON)
		r.Get("/", handler.handlerRoot)
	})
	return
}

func (t *DefaultHandler) checkCompress(r *http.Request) (io.Reader, error) {
	if checkContentEncodingGZIP(r) {
		if b, err := gzip.Decompress(&gzip.DecompressParams{
			Reader: r.Body,
		}); err != nil {
			return nil, fmt.Errorf("decompress failed: %w", err)
		} else {
			return b, nil
		}
	} else {
		return r.Body, nil
	}
}
