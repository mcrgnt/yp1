package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mcrgnt/yp1/internal/storage"

	"github.com/microgiantya/logger"
)

var (
	htmlHeader = `<!DOCTYPE html><html><head><title>Metrics</title></head><body>`
	htmlFooter = `</body></html>`
)

type DefaultHandler struct {
	storage storage.MemStorage
	ctx     *logger.Logger
	R       *chi.Mux
}

type NewDefaultHandlerParams struct {
	Storage storage.MemStorage
}

func (t *DefaultHandler) writeResponse(w http.ResponseWriter, r *http.Request, statusHeader int, err error) {
	t.ctx.LogInformational(fmt.Sprintf("new request: method: %s, path: %s", r.Method, r.URL.Path))
	if err != nil {
		t.ctx.LogError(err)
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

	updateParams := &storage.StorageParams{
		Type: chi.URLParam(r, "type"),
		Name: chi.URLParam(r, "name"),
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

	err = t.storage.GetByType(updateParams)
	if err != nil {
		statusHeader = http.StatusNotFound
		return
	}
	_, _ = fmt.Fprintf(w, "%v", updateParams.Value)
}

func (t *DefaultHandler) handlerRoot(w http.ResponseWriter, r *http.Request) {
	var (
		statusHeader = 200
		err          error
	)
	defer t.writeResponse(w, r, statusHeader, err)
	_, _ = w.Write([]byte(htmlHeader + t.storage.GetAll() + htmlFooter))
}

func NewDefaultHandler(ctx context.Context, params *NewDefaultHandlerParams) (handler *DefaultHandler) {
	handler = &DefaultHandler{
		storage: params.Storage,
		ctx: logger.NewLoggerContext(ctx, &logger.LoggerInitParams{
			Severity:       logSeverity,
			UniqueIDPrefix: "hdl",
			Version:        "v-",
		}),
		R: chi.NewRouter(),
	}

	handler.R.Post("/update/{type}/{name}/{value}", handler.handlerUpdate)
	handler.R.Get("/value/{type}/{name}", handler.handlerValue)
	handler.R.Get("/", handler.handlerRoot)

	return
}
