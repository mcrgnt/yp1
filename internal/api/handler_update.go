package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mcrgnt/yp1/internal/store/models"
)

func (t *DefaultHandler) handlerUpdate(w http.ResponseWriter, r *http.Request) {
	var (
		err          error
		statusHeader = http.StatusOK
	)

	defer func() {
		if err != nil {
			switch {
			case errors.Is(err, models.ErrEmptyMetricName):
				statusHeader = http.StatusNotFound
			default:
				statusHeader = http.StatusBadRequest
			}
		}
		w.WriteHeader(statusHeader)
	}()
	storageParams := &models.StorageParams{
		Type:  chi.URLParam(r, "type"),
		Name:  chi.URLParam(r, "name"),
		Value: chi.URLParam(r, "value"),
	}
	err = t.storage.MetricSet(storageParams)
}

func (t *DefaultHandler) handlerUpdateJSON(w http.ResponseWriter, r *http.Request) {
	var (
		err           error
		statusHeader  = http.StatusOK
		storageParams = &models.StorageParams{}
		returnBody    []byte
	)

	checkAcceptEncodingGZIP(w, r)
	w.Header().Set(contentType, applicationJSON)

	defer func() {
		if len(returnBody) == 0 {
			err = models.ErrMetricNotFound
		}
		if err != nil {
			switch {
			case errors.Is(err, models.ErrEmptyMetricName):
				statusHeader = http.StatusNotFound
			default:
				statusHeader = http.StatusBadRequest
			}
		}
		w.WriteHeader(statusHeader)
		_, _ = w.Write(returnBody)
	}()

	switch r.Header.Get(contentType) {
	case applicationJSON:
		if b, err := t.checkCompress(r); err != nil {
			return
		} else {
			if err = json.NewDecoder(b).Decode(storageParams); err != nil {
				return
			}
		}
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
	if returnBody, err = json.Marshal(storageParams); err != nil {
		return
	}
}
