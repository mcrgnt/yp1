package metric

import (
	"strconv"
	"testing"

	"github.com/mcrgnt/yp1/internal/common"
	"github.com/stretchr/testify/assert"
)

func TestGauge(t *testing.T) {
	tests := []struct {
		params              *NewGaugeParams
		expectedType        string
		expectedValue       float64
		expectedString      string
		expectedResetString string
		name                string
	}{
		{
			name:                "test_a",
			params:              &NewGaugeParams{Value: "0"},
			expectedType:        common.TypeMetricGauge,
			expectedValue:       float64(0),
			expectedString:      "0",
			expectedResetString: "0",
		},
		{
			name:                "test_a",
			params:              &NewGaugeParams{Value: "10"},
			expectedType:        common.TypeMetricGauge,
			expectedValue:       float64(10),
			expectedString:      "10",
			expectedResetString: "0",
		},
	}

	for i, tt := range tests {
		t.Run(tt.name+"_"+strconv.Itoa(i), func(t *testing.T) {
			actual, actualErr := NewGauge(tt.params)
			assert.Equal(t, nil, actualErr)
			assert.Equal(t, tt.expectedType, actual.Type())
			assert.Equal(t, tt.expectedValue, actual.Value)
			assert.Equal(t, tt.expectedString, actual.String())
			actual.Reset()
			assert.Equal(t, tt.expectedResetString, actual.String())
		})
	}
}
