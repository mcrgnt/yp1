package reporter

import (
	"fmt"
	"io"
	"net/http"

	"github.com/microgiantya/logger"
)

type Reporter struct {
	ctx *logger.Logger
}

type ReportParams struct {
	Url string
}

func (t *Reporter) Report(params *ReportParams) {
	t.ctx.LogInformational(fmt.Sprintf("new report with params: %v", *params))
	resp, err := http.Post(params.Url, "", nil)
	if err != nil {
		t.ctx.LogError(fmt.Sprintf("report response: %v", err))
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.ctx.LogError(fmt.Sprintf("report response: %v", err))
		return
	}
	t.ctx.LogInformational(fmt.Sprintf("report response: %v", string(bodyBytes)))
}
