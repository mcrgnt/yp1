package reporter

import (
	"fmt"
	"io"
	"net/http"

	"github.com/microgiantya/logger"
)

const (
	logSeverity = 7
)

type Reporter struct {
	ctx *logger.Logger
}

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
	reporter := &Reporter{
		ctx: logger.NewLogger(&logger.LoggerInitParams{
			UniqueIDPrefix: "rpt",
			Version:        "v-",
			Severity:       logSeverity,
		}),
	}
	defer reporter.ctx.Close()

	reporter.ctx.LogInformational(fmt.Sprintf("new report with params: %v", *params))
	response, err := reporter.report(params)
	if err != nil {
		reporter.ctx.LogError(fmt.Sprintf("report response: %v", err))
		return
	}
	reporter.ctx.LogInformational(fmt.Sprintf("report response: %v", response))
}
