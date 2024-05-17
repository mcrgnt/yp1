package metric

import (
	"strconv"
	"testing"

	"github.com/mcrgnt/yp1/internal/common"
	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {
	tests := []struct {
		params              *NewCounterParams
		expectedType        string
		expectedValue       int64
		expectedString      string
		expectedResetString string
		name                string
	}{
		{
			name:                "test_a",
			params:              &NewCounterParams{Value: "0"},
			expectedType:        common.TypeMetricCounter,
			expectedValue:       int64(0),
			expectedString:      "0",
			expectedResetString: "0",
		},
		{
			name:                "test_a",
			params:              &NewCounterParams{Value: "10"},
			expectedType:        common.TypeMetricCounter,
			expectedValue:       int64(10),
			expectedString:      "10",
			expectedResetString: "0",
		},
	}

	for i, tt := range tests {
		t.Run(tt.name+"_"+strconv.Itoa(i), func(t *testing.T) {
			actual, actualErr := NewCounter(tt.params)
			assert.Equal(t, nil, actualErr)
			assert.Equal(t, tt.expectedType, actual.Type())
			assert.Equal(t, tt.expectedValue, actual.Value)
			assert.Equal(t, tt.expectedString, actual.String())
			actual.Reset()
			assert.Equal(t, tt.expectedResetString, actual.String())
		})
	}
}
