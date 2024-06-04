package reporter

import (
	"bytes"
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
	resp, err := http.Post(params.URL, "application/json", bytes.NewReader(params.Body))
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
