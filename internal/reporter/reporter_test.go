package reporter

import (
	"testing"

	"github.com/microgiantya/logger"
)

var (
	loggerInitParams = &logger.LoggerInitParams{}
)

func TestReporter_report(t *testing.T) {
	type fields struct {
		ctx *logger.Logger
	}
	type args struct {
		params *ReportParams
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantResponse string
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name: "test0",
			fields: fields{
				ctx: logger.NewLogger(loggerInitParams),
			},
			args: args{
				params: &ReportParams{
					URL: "http://localhost:8080/update/1/2/3",
				},
			},
			wantResponse: ``,
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Reporter{}
			gotResponse, err := tr.report(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reporter.report() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponse != tt.wantResponse {
				t.Errorf("Reporter.report() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
