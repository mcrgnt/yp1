package storage

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/mcrgnt/yp1/internal/common"
	"github.com/mcrgnt/yp1/internal/storage/internal/metric"
	"github.com/stretchr/testify/assert"
)

func Test_Equal_NewMemStorage(t *testing.T) {
	tests := []struct {
		want *MemStorage
		name string
	}{
		{
			name: "test0",
			want: &MemStorage{
				Metrics: map[string]metric.Metric{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewMemStorage()
			assert.Equal(t, reflect.TypeOf(tt.want), reflect.TypeOf(actual))
			assert.Equal(t, tt.want.Metrics, actual.Metrics)
		})
	}
}
func Test_NotEqual_NewMemStorage(t *testing.T) {
	tests := []struct {
		want *MemStorage
		name string
	}{
		{
			name: "test0",
			want: &MemStorage{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewMemStorage()
			assert.NotEqual(t, nil, actual)
			assert.NotEqual(t, tt.want.Metrics, actual.Metrics)
		})
	}
}

func TestMemStorage_MetricSetGaugeEqual(t *testing.T) {
	tests := []struct {
		expected string
		params   *StorageParams
		name     string
	}{
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: float64(-1)}, expected: "-1"},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: float64(-1.0)}, expected: "-1"},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: float64(-1.0001)}, expected: "-1.0001"},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: float64(0)}, expected: "0"},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: float64(0.0)}, expected: "0"},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: float64(0.0001)}, expected: "0.0001"},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: float64(1)}, expected: "1"},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: float64(1.0)}, expected: "1"},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: float64(1.0001)}, expected: "1.0001"},

		{name: "test_b", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: uint64(0)}, expected: "0"},
		{name: "test_b", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: uint64(0.0)}, expected: "0"},
		{name: "test_b", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: uint64(1)}, expected: "1"},
		{name: "test_b", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: uint64(1.0)}, expected: "1"},

		{name: "test_c", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: uint32(0)}, expected: "0"},
		{name: "test_c", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: uint32(0.0)}, expected: "0"},
		{name: "test_c", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: uint32(1)}, expected: "1"},
		{name: "test_c", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: uint32(1.0)}, expected: "1"},

		{name: "test_d", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: "-1"}, expected: "-1"},
		{name: "test_d", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: "-1.0"}, expected: "-1"},
		{name: "test_d", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: "-1.0001"}, expected: "-1.0001"},
		{name: "test_d", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: "0"}, expected: "0"},
		{name: "test_d", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: "0.0"}, expected: "0"},
		{name: "test_d", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: "0.0001"}, expected: "0.0001"},
		{name: "test_d", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: "1"}, expected: "1"},
		{name: "test_d", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: "1.0"}, expected: "1"},
		{name: "test_d", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: "1.0001"}, expected: "1.0001"},
	}
	for i, tt := range tests {
		t.Run(tt.name+"_"+strconv.Itoa(i), func(t *testing.T) {
			_storage := NewMemStorage()
			actualErr := _storage.MetricSet(tt.params)
			assert.Equal(t, nil, actualErr)
			actualErr = _storage.GetMetricString(tt.params)
			assert.Equal(t, nil, actualErr)
			assert.Equal(t, tt.expected, tt.params.String)
		})
	}
}

func TestMemStorage_MetricSetCounterEqual(t *testing.T) {
	tests := []struct {
		expected string
		params   *StorageParams
		name     string
	}{
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricCounter, Name: "test", Value: int64(0)}, expected: "0"},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricCounter, Name: "test", Value: int64(0.0)}, expected: "0"},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricCounter, Name: "test", Value: int64(1)}, expected: "1"},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricCounter, Name: "test", Value: int64(1.0)}, expected: "1"},

		{name: "test_b", params: &StorageParams{Type: common.TypeMetricCounter, Name: "test", Value: "0"}, expected: "0"},
		{name: "test_b", params: &StorageParams{Type: common.TypeMetricCounter, Name: "test", Value: "1"}, expected: "1"},
	}
	for i, tt := range tests {
		t.Run(tt.name+"_"+strconv.Itoa(i), func(t *testing.T) {
			_storage := NewMemStorage()
			actualErr := _storage.MetricSet(tt.params)
			assert.Equal(t, nil, actualErr)
			actualErr = _storage.GetMetricString(tt.params)
			assert.Equal(t, nil, actualErr)
			assert.Equal(t, tt.expected, tt.params.String)
		})
	}
}

func TestMemStorage_MetricSetGaugeErr(t *testing.T) {
	tests := []struct {
		params      *StorageParams
		name        string
		expectedErr error
	}{
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: int32(-1)}, expectedErr: common.ErrIncompatibleMetricValueType},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: int32(0)}, expectedErr: common.ErrIncompatibleMetricValueType},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: int32(1)}, expectedErr: common.ErrIncompatibleMetricValueType},

		{name: "test_b", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: ""}, expectedErr: strconv.ErrSyntax},
		{name: "test_b", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: "1.79769313486231570814527423731704356798070e+309"}, expectedErr: strconv.ErrRange},
		{name: "test_b", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: "1.000000o1"}, expectedErr: strconv.ErrSyntax},
	}
	for i, tt := range tests {
		t.Run(tt.name+"_"+strconv.Itoa(i), func(t *testing.T) {
			_storage := NewMemStorage()
			actualErr := _storage.MetricSet(tt.params)
			assert.ErrorIs(t, actualErr, tt.expectedErr)
		})
	}
}

func TestMemStorage_MetricSetCounterErr(t *testing.T) {
	tests := []struct {
		params      *StorageParams
		name        string
		expectedErr error
	}{
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricCounter, Name: "test", Value: int32(-1)}, expectedErr: common.ErrIncompatibleMetricValueType},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricCounter, Name: "test", Value: int32(0)}, expectedErr: common.ErrIncompatibleMetricValueType},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricCounter, Name: "test", Value: int32(1)}, expectedErr: common.ErrIncompatibleMetricValueType},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricCounter, Name: "test", Value: int64(-1)}, expectedErr: common.ErrIncompatibleMetricValue},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricCounter, Name: "test", Value: float64(-1)}, expectedErr: common.ErrIncompatibleMetricValueType},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricCounter, Name: "test", Value: float64(-1.00001)}, expectedErr: common.ErrIncompatibleMetricValueType},

		{name: "test_b", params: &StorageParams{Type: common.TypeMetricCounter, Name: "test", Value: ""}, expectedErr: strconv.ErrSyntax},
		{name: "test_b", params: &StorageParams{Type: common.TypeMetricCounter, Name: "test", Value: "1.79769313486231570814527423731704356798070e+309"}, expectedErr: strconv.ErrSyntax},
		{name: "test_b", params: &StorageParams{Type: common.TypeMetricCounter, Name: "test", Value: "1.000000o1"}, expectedErr: strconv.ErrSyntax},
		{name: "test_b", params: &StorageParams{Type: common.TypeMetricCounter, Name: "test", Value: "-1"}, expectedErr: common.ErrIncompatibleMetricValue},
	}
	for i, tt := range tests {
		t.Run(tt.name+"_"+strconv.Itoa(i), func(t *testing.T) {
			_storage := NewMemStorage()
			actualErr := _storage.MetricSet(tt.params)
			assert.ErrorIs(t, actualErr, tt.expectedErr)
		})
	}
}

func TestMemStorage_MetricSetWrongType(t *testing.T) {
	tests := []struct {
		params      *StorageParams
		name        string
		expectedErr error
	}{
		{name: "test_a", params: &StorageParams{Type: "", Name: "test", Value: int64(0)}, expectedErr: common.ErrNotImplementedMetricType},
		{name: "test_a", params: &StorageParams{Type: "empty", Name: "test", Value: int64(0)}, expectedErr: common.ErrNotImplementedMetricType},
	}
	for i, tt := range tests {
		t.Run(tt.name+"_"+strconv.Itoa(i), func(t *testing.T) {
			_storage := NewMemStorage()
			actualErr := _storage.MetricSet(tt.params)
			assert.ErrorIs(t, actualErr, tt.expectedErr)
		})
	}
}

func TestMemStorage_MetricSetEmptyName(t *testing.T) {
	tests := []struct {
		params      *StorageParams
		name        string
		expectedErr error
	}{
		{name: "test_a", params: &StorageParams{Type: "", Name: "", Value: int64(0)}, expectedErr: common.ErrEmptyMetricName},
	}
	for i, tt := range tests {
		t.Run(tt.name+"_"+strconv.Itoa(i), func(t *testing.T) {
			_storage := NewMemStorage()
			actualErr := _storage.MetricSet(tt.params)
			assert.ErrorIs(t, actualErr, tt.expectedErr)
		})
	}
}

func TestMemStorage_MetricGaugeSequentialSet(t *testing.T) {
	tests := []struct {
		params   *StorageParams
		name     string
		expected string
	}{
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: float64(0)}, expected: "0"},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: float64(3)}, expected: "3"},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: float64(6)}, expected: "6"},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: float64(9)}, expected: "9"},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: float64(0)}, expected: "0"},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: float64(-3)}, expected: "-3"},
	}

	_storage := NewMemStorage()

	for i, tt := range tests {
		t.Run(tt.name+"_"+strconv.Itoa(i), func(t *testing.T) {
			actualErr := _storage.MetricSet(tt.params)
			assert.Equal(t, nil, actualErr)
			actualErr = _storage.GetMetricString(tt.params)
			assert.Equal(t, nil, actualErr)
			assert.Equal(t, tt.expected, tt.params.String)
		})
	}
}

func TestMemStorage_MetricCounterSequentialSet(t *testing.T) {
	tests := []struct {
		params   *StorageParams
		name     string
		expected string
	}{
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricCounter, Name: "test", Value: int64(0)}, expected: "0"},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricCounter, Name: "test", Value: int64(3)}, expected: "3"},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricCounter, Name: "test", Value: int64(3)}, expected: "6"},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricCounter, Name: "test", Value: int64(3)}, expected: "9"},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricCounter, Name: "test", Value: int64(3)}, expected: "12"},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricCounter, Name: "test", Value: int64(3)}, expected: "15"},
	}

	_storage := NewMemStorage()

	for i, tt := range tests {
		t.Run(tt.name+"_"+strconv.Itoa(i), func(t *testing.T) {
			actualErr := _storage.MetricSet(tt.params)
			assert.Equal(t, nil, actualErr)
			actualErr = _storage.GetMetricString(tt.params)
			assert.Equal(t, nil, actualErr)
			assert.Equal(t, tt.expected, tt.params.String)
		})
	}
}

func TestMemStorage_MetricSequentialSet(t *testing.T) {
	tests := []struct {
		params      *StorageParams
		name        string
		expected    string
		expectedErr error
	}{
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: float64(0)}, expected: "0"},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricCounter, Name: "test", Value: int64(3)}, expected: "3"},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: float64(3)}, expected: "3"},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricCounter, Name: "test", Value: int64(3)}, expected: "6"},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test", Value: float64(-1)}, expected: "-1"},
		{name: "test_a", params: &StorageParams{Type: common.TypeMetricCounter, Name: "test", Value: int64(3)}, expected: "9"},

		{name: "test_b", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test1", Value: float64(888)}, expected: "888"},
		{name: "test_b", params: &StorageParams{Type: common.TypeMetricGauge, Name: "test1", Value: int64(-1)}, expected: "888", expectedErr: common.ErrIncompatibleMetricValueType},

		{name: "test_c", params: &StorageParams{Type: common.TypeMetricCounter, Name: "test2", Value: int64(999)}, expected: "999"},
		{name: "test_c", params: &StorageParams{Type: common.TypeMetricCounter, Name: "test2", Value: "-1"}, expected: "999", expectedErr: common.ErrIncompatibleMetricValue},
		{name: "test_c", params: &StorageParams{Type: common.TypeMetricCounter, Name: "test2", Value: int64(999)}, expected: "1998"},
		{name: "test_c", params: &StorageParams{Type: common.TypeMetricCounter, Name: "test2", Value: "-1"}, expected: "1998", expectedErr: common.ErrIncompatibleMetricValue},
	}

	_storage := NewMemStorage()

	for i, tt := range tests {
		t.Run(tt.name+"_"+strconv.Itoa(i), func(t *testing.T) {
			actualErr := _storage.MetricSet(tt.params)
			if tt.expectedErr == nil {
				assert.Equal(t, nil, actualErr)
			} else {
				assert.ErrorIs(t, actualErr, tt.expectedErr)
			}
			actualErr = _storage.GetMetricString(tt.params)
			assert.Equal(t, nil, actualErr)
			assert.Equal(t, tt.expected, tt.params.String)
		})
	}
}
