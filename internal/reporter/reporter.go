package reporter

import (
	"fmt"
	"io"
	"net/http"

	"github.com/mcrgnt/yp1/internal/common"
	"github.com/mcrgnt/yp1/internal/compress/gzip"
)

type Reporter struct{}

type ReportParams struct {
	URL  string
	Body []byte
}

func (t *Reporter) report(params *ReportParams) (string, error) {
	if buf, err := gzip.Compress(params.Body); err != nil {
		return "", fmt.Errorf("compress failed: %w", err)
	} else {
		if req, err := http.NewRequest("POST", params.URL, buf); err != nil {
			return "", fmt.Errorf("new request failed: %w", err)
		} else {
			req.Header.Set(common.ContentType, common.ApplicationJSON)
			req.Header.Set(common.ContentEncoding, common.GZip)
			req.Header.Set(common.AcceptEncoding, common.GZip)

			if resp, err := http.DefaultClient.Do(req); err != nil {
				return "", fmt.Errorf("report response failed: %w", err)
			} else {
				defer func() {
					if err := resp.Body.Close(); err != nil {
						fmt.Println(err)
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
}

func Report(params *ReportParams) (err error) {
	reporter := &Reporter{}
	_, err = reporter.report(params)
	return
}
