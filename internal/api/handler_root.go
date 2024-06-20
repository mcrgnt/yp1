package api

import (
	"net/http"

	"github.com/mcrgnt/yp1/internal/common"
)

func (t *DefaultHandler) handlerRoot(w http.ResponseWriter, r *http.Request) {
	common.CheckAcceptEncodingGZIP(w, r)
	w.Header().Set(common.ContentType, common.TextHTML)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(htmlHeader + t.storage.GetMetricAll() + htmlFooter))
}
