package reporter

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/mcrgnt/yp1/internal/compress/gzip"
	"github.com/rs/zerolog"
)

const (
	contentType     = "Content-Type"
	contentEncoding = "Content-Encoding"
	acceptEncoding  = "Accept-Encoding"

	applicationJSON = "application/json"
	textHTML        = "text/html"
	gZip            = "gzip"
)

type Reporter struct{}

type ReportParams struct {
	Logger *zerolog.Logger
	URL    string
	Body   []byte
}

func (t *Reporter) report(params *ReportParams) (data string, err error) {
	var (
		buf       io.Reader
		req       *http.Request
		resp      *http.Response
		bodyBytes []byte
	)

	if buf, err = gzip.Compress(params.Body); err != nil {
		err = fmt.Errorf("compress failed: %w", err)
		return
	}

	if req, err = http.NewRequest(http.MethodPost, params.URL, buf); err != nil {
		err = fmt.Errorf("new request failed: %w", err)
		return
	}

	req.Header.Set(contentType, applicationJSON)
	req.Header.Set(contentEncoding, gZip)
	req.Header.Set(acceptEncoding, gZip)

	if resp, err = http.DefaultClient.Do(req); err != nil {
		err = fmt.Errorf("request do: %w", err)
		return
	}
	defer func() {
		if e := resp.Body.Close(); e != nil {
			data = ""
			err = errors.Join(err, fmt.Errorf("body close failed: %w", e))
		}
	}()

	if bodyBytes, err = io.ReadAll(resp.Body); err != nil {
		err = fmt.Errorf("read all failed: %w", err)
		return
	}
	data = string(bodyBytes)
	return
}

func Report(params *ReportParams) (err error) {
	reporter := &Reporter{}
	_, err = reporter.report(params)
	return
}
