package api

import (
	"net/http"

	"github.com/mcrgnt/yp1/internal/db"
)

func (t *DefaultHandler) handlerPing(w http.ResponseWriter, r *http.Request) {
	if err := db.Ping(t.databaseDSN); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
