package metric

import (
	"strconv"
	"testing"

	"github.com/mcrgnt/yp1/internal/common"
	"github.com/stretchr/testify/assert"
)

func TestGuageCounter(t *testing.T) {
	tests := []struct {
		params              any
		expectedType        string
		expectedValue       any
		expectedString      string
		expectedResetString string
		name                string
	}{
		{
			name:                "test_a",
			params:              &NewGaugeParams{Val: "0"},
			expectedType:        common.TypeMetricGauge,
			expectedString:      "0",
			expectedResetString: "0",
		},
		{
			name:                "test_a",
			params:              &NewGaugeParams{Val: "10"},
			expectedType:        common.TypeMetricGauge,
			expectedString:      "10",
			expectedResetString: "0",
		},
		{
			name:                "test_b",
			params:              &NewCounterParams{Val: "0"},
			expectedType:        common.TypeMetricCounter,
			expectedString:      "0",
			expectedResetString: "0",
		},
		{
			name:                "test_b",
			params:              &NewCounterParams{Val: "10"},
			expectedType:        common.TypeMetricCounter,
			expectedString:      "10",
			expectedResetString: "0",
		},
	}

	for i, tt := range tests {
		t.Run(tt.name+"_"+strconv.Itoa(i), func(t *testing.T) {
			var (
				actual    Metric
				actualErr error
			)
			switch tt.expectedType {
			case common.TypeMetricCounter:
				actual, actualErr = NewCounter(tt.params.(*NewCounterParams))
			default:
				actual, actualErr = NewGauge(tt.params.(*NewGaugeParams))
			}
			assert.Equal(t, nil, actualErr)
			assert.Equal(t, tt.expectedType, actual.Type())
			assert.Equal(t, tt.expectedString, actual.String())
			actual.Reset()
			assert.Equal(t, tt.expectedResetString, actual.String())
		})
	}
}
