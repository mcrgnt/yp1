package common

import (
	"net/http"
	"strings"
)

func CheckContentEncodingGZIP(r *http.Request) bool {
	return strings.Contains(r.Header.Get(ContentEncoding), strings.ToLower(GZip))
}

func CheckAcceptEncodingGZIP(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.Header.Get(ContentEncoding), strings.ToLower(GZip)) {
		w.Header().Set(ContentEncoding, GZip)
	}
}
