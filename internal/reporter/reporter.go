package reporter

import (
	"fmt"
	"io"
	"net/http"
)

type Reporter struct{}

type ReportParams struct {
	URL string
}

func (t *Reporter) report(params *ReportParams) (response string, err error) {
	resp, err := http.Post(params.URL, "text/plain", nil)
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

func Report(params *ReportParams) {
	reporter := &Reporter{}
	_, _ = reporter.report(params)
}
