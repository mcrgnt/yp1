package api

import (
	"net/http"
)

func (t *DefaultHandler) handlerRoot(w http.ResponseWriter, r *http.Request) {
	checkAcceptEncodingGZIP(w, r)
	w.Header().Set(contentType, textHTML)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(htmlHeader + t.storage.GetMetricAll() + htmlFooter)); err != nil {
		t.logger.Error().Msgf("write failed: %s", err)
	}
}
