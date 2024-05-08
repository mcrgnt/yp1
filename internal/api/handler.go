package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/mcrgnt/yp1/internal/storage"

	"github.com/microgiantya/logger"
)

const (
	pathUpdateChunksLen = 4
)

type DefaultHandler struct {
	storage storage.MemStorage
	ctx     *logger.Logger
}

func (t *DefaultHandler) handleUpdate(r *http.Request, pathChunks []string) (statusHeader int, err error) {
	if r.Method != http.MethodPost {
		statusHeader = http.StatusBadRequest
		err = fmt.Errorf("http method must be POST, actual: %s", r.Method)
		return
	}

	if len(pathChunks) < pathUpdateChunksLen {
		statusHeader = http.StatusBadRequest
		err = fmt.Errorf("path have less data then required %d, actual: %d", pathUpdateChunksLen, len(pathChunks))
		return
	}

	updateParams := &storage.StorageParams{
		Type:  pathChunks[1],
		Name:  pathChunks[2],
		Value: pathChunks[3],
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
	return
}

func (t *DefaultHandler) parsePath(r *http.Request) []string {
	return strings.Split(strings.Trim(r.URL.Path, "/"), "/")
}

func (t *DefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		statusHeader = 200
		err          error
	)
	defer func() {
		t.ctx.LogInformational(fmt.Sprintf("new request: method: %s, path: %s", r.Method, r.URL.Path))
		if err != nil {
			t.ctx.LogError(err)
		}

		w.WriteHeader(statusHeader)
	}()

	pathChunks := t.parsePath(r)

	switch pathChunks[0] {
	case "update":
		statusHeader, err = t.handleUpdate(r, pathChunks)
	default:
		statusHeader = http.StatusNotFound
		err = fmt.Errorf("unknown path: %s", pathChunks[0])
	}

}
