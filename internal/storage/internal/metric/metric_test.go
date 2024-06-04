package metric

import (
	"strconv"
	"testing"

	"github.com/mcrgnt/yp1/internal/common"
	"github.com/stretchr/testify/assert"
)

func TestNewMetric(t *testing.T) {
	tests := []struct {
		expected    Metric
		expectedErr error
		params      *NewMetricParams
		name        string
	}{
		{
			name:     "test_",
			params:   &NewMetricParams{Type: common.TypeMetricGauge, Value: float64(1)},
			expected: &Gauge{Value: 1}},
		{
			name:     "test_",
			params:   &NewMetricParams{Type: common.TypeMetricCounter, Value: int64(1)},
			expected: &Counter{Value: 1}},
		{
			name:        "test_",
			params:      &NewMetricParams{Type: "", Value: int64(1)},
			expected:    nil,
			expectedErr: common.ErrNotImplementedMetricType,
		},
	}

	for i, tt := range tests {
		t.Run(tt.name+"_"+strconv.Itoa(i), func(t *testing.T) {
			actual, actualErr := NewMetric(tt.params)
			if tt.expectedErr == nil {
				assert.Equal(t, tt.expected, actual)
			} else {
				assert.ErrorIs(t, actualErr, tt.expectedErr)
			}
		})
	}
}
