package reporter

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
)

type Reporter struct{}

type ReportParams struct {
	URL  string
	Body []byte
}

func (t *Reporter) report(params *ReportParams) (response string, err error) {
	var buf bytes.Buffer
	g := gzip.NewWriter(&buf)

	if _, err = g.Write(params.Body); err != nil {
		return
	}

	if err = g.Close(); err != nil {
		return
	}

	var (
		req *http.Request
	)
	if req, err = http.NewRequest("POST", params.URL, &buf); err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("Content-Encoding", "gzip")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		err = fmt.Errorf("report response: %w", err)
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("report response: %w", err)
		return
	}

	response = string(bodyBytes)
	return
}

func Report(params *ReportParams) (err error) {
	reporter := &Reporter{}
	_, err = reporter.report(params)
	return
}
