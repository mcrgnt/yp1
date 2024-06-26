package reporter

import (
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

func (t *Reporter) report(params *ReportParams) (string, error) {
	if buf, err := gzip.Compress(params.Body); err != nil {
		return "", fmt.Errorf("compress failed: %w", err)
	} else if req, err := http.NewRequest(http.MethodPost, params.URL, buf); err != nil {
		return "", fmt.Errorf("new request failed: %w", err)
	} else {
		req.Header.Set(contentType, applicationJSON)
		req.Header.Set(contentEncoding, gZip)
		req.Header.Set(acceptEncoding, gZip)

		if resp, err := http.DefaultClient.Do(req); err != nil {
			return "", fmt.Errorf("report response failed: %w", err)
		} else {
			defer func() {
				if err := resp.Body.Close(); err != nil {
					params.Logger.Error().Msg(err.Error())
				}
			}()

			if bodyBytes, err := io.ReadAll(resp.Body); err != nil {
				return "", fmt.Errorf("report response failed: %w", err)
			} else {
				return string(bodyBytes), nil
			}
		}
	}
}

func Report(params *ReportParams) (err error) {
	reporter := &Reporter{}
	_, err = reporter.report(params)
	return
}
