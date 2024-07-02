package memory

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/mcrgnt/yp1/internal/store/models"
	"github.com/stretchr/testify/assert"
)

func Test_Equal_NewMemoryStorage(t *testing.T) {
	tests := []struct {
		want *MemoryStorage
		name string
	}{
		{
			name: "test0",
			want: &MemoryStorage{
				Metrics: map[string]models.Metric{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewMemoryStorage()
			assert.Equal(t, reflect.TypeOf(tt.want), reflect.TypeOf(actual))
			assert.Equal(t, tt.want.Metrics, actual.Metrics)
		})
	}
}
func Test_NotEqual_NewMemoryStorage(t *testing.T) {
	tests := []struct {
		want *MemoryStorage
		name string
	}{
		{
			name: "test0",
			want: &MemoryStorage{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewMemoryStorage()
			assert.NotEqual(t, nil, actual)
			assert.NotEqual(t, tt.want.Metrics, actual.Metrics)
		})
	}
}

func TestMemoryStorage_MetricSetGaugeEqual(t *testing.T) {
	tests := []struct {
		expected string
		params   *models.StorageParams
		name     string
	}{
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: float64(-1)}, expected: "-1"},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: float64(-1.0)}, expected: "-1"},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: float64(-1.0001)}, expected: "-1.0001"},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: float64(0)}, expected: "0"},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: float64(0.0)}, expected: "0"},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: float64(0.0001)}, expected: "0.0001"},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: float64(1)}, expected: "1"},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: float64(1.0)}, expected: "1"},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: float64(1.0001)}, expected: "1.0001"},

		{name: "test_b", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: uint64(0)}, expected: "0"},
		{name: "test_b", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: uint64(0.0)}, expected: "0"},
		{name: "test_b", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: uint64(1)}, expected: "1"},
		{name: "test_b", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: uint64(1.0)}, expected: "1"},

		{name: "test_c", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: uint32(0)}, expected: "0"},
		{name: "test_c", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: uint32(0.0)}, expected: "0"},
		{name: "test_c", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: uint32(1)}, expected: "1"},
		{name: "test_c", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: uint32(1.0)}, expected: "1"},

		{name: "test_d", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: "-1"}, expected: "-1"},
		{name: "test_d", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: "-1.0"}, expected: "-1"},
		{name: "test_d", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: "-1.0001"}, expected: "-1.0001"},
		{name: "test_d", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: "0"}, expected: "0"},
		{name: "test_d", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: "0.0"}, expected: "0"},
		{name: "test_d", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: "0.0001"}, expected: "0.0001"},
		{name: "test_d", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: "1"}, expected: "1"},
		{name: "test_d", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: "1.0"}, expected: "1"},
		{name: "test_d", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: "1.0001"}, expected: "1.0001"},
	}
	for i, tt := range tests {
		t.Run(tt.name+"_"+strconv.Itoa(i), func(t *testing.T) {
			_storage := NewMemoryStorage()
			actualErr := _storage.MetricSet(tt.params)
			assert.Equal(t, nil, actualErr)
			actualErr = _storage.GetMetricString(tt.params)
			assert.Equal(t, nil, actualErr)
			assert.Equal(t, tt.expected, tt.params.String)
		})
	}
}

func TestMemoryStorage_MetricSetCounterEqual(t *testing.T) {
	tests := []struct {
		expected string
		params   *models.StorageParams
		name     string
	}{
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricCounter,
			Name: "test", Value: int64(0)}, expected: "0"},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricCounter,
			Name: "test", Value: int64(0.0)}, expected: "0"},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricCounter,
			Name: "test", Value: int64(1)}, expected: "1"},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricCounter,
			Name: "test", Value: int64(1.0)}, expected: "1"},

		{name: "test_b", params: &models.StorageParams{Type: models.TypeMetricCounter,
			Name: "test", Value: "0"}, expected: "0"},
		{name: "test_b", params: &models.StorageParams{Type: models.TypeMetricCounter,
			Name: "test", Value: "1"}, expected: "1"},
	}
	for i, tt := range tests {
		t.Run(tt.name+"_"+strconv.Itoa(i), func(t *testing.T) {
			_storage := NewMemoryStorage()
			actualErr := _storage.MetricSet(tt.params)
			assert.Equal(t, nil, actualErr)
			actualErr = _storage.GetMetricString(tt.params)
			assert.Equal(t, nil, actualErr)
			assert.Equal(t, tt.expected, tt.params.String)
		})
	}
}

func TestMemoryStorage_MetricSetGaugeErr(t *testing.T) {
	tests := []struct {
		expectedErr error
		params      *models.StorageParams
		name        string
	}{
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: int32(-1)}, expectedErr: models.ErrIncompatibleMetricValueType},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: int32(0)}, expectedErr: models.ErrIncompatibleMetricValueType},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: int32(1)}, expectedErr: models.ErrIncompatibleMetricValueType},

		{name: "test_b", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: ""}, expectedErr: strconv.ErrSyntax},
		{name: "test_b", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: "1.79769313486231570814527423731704356798070e+309"}, expectedErr: strconv.ErrRange},
		{name: "test_b", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: "1.000000o1"}, expectedErr: strconv.ErrSyntax},
	}
	for i, tt := range tests {
		t.Run(tt.name+"_"+strconv.Itoa(i), func(t *testing.T) {
			_storage := NewMemoryStorage()
			actualErr := _storage.MetricSet(tt.params)
			assert.ErrorIs(t, actualErr, tt.expectedErr)
		})
	}
}

func TestMemoryStorage_MetricSetCounterErr(t *testing.T) {
	tests := []struct {
		expectedErr error
		params      *models.StorageParams
		name        string
	}{
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricCounter,
			Name: "test", Value: int32(-1)}, expectedErr: models.ErrIncompatibleMetricValueType},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricCounter,
			Name: "test", Value: int32(0)}, expectedErr: models.ErrIncompatibleMetricValueType},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricCounter,
			Name: "test", Value: int32(1)}, expectedErr: models.ErrIncompatibleMetricValueType},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricCounter,
			Name: "test", Value: int64(-1)}, expectedErr: models.ErrIncompatibleMetricValue},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricCounter,
			Name: "test", Value: float64(-1)}, expectedErr: models.ErrIncompatibleMetricValue},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricCounter,
			Name: "test", Value: float64(-1.00001)}, expectedErr: models.ErrIncompatibleMetricValue},

		{name: "test_b", params: &models.StorageParams{Type: models.TypeMetricCounter,
			Name: "test", Value: ""}, expectedErr: strconv.ErrSyntax},
		{name: "test_b", params: &models.StorageParams{Type: models.TypeMetricCounter,
			Name: "test", Value: "1.79769313486231570814527423731704356798070e+309"}, expectedErr: strconv.ErrSyntax},
		{name: "test_b", params: &models.StorageParams{Type: models.TypeMetricCounter,
			Name: "test", Value: "1.000000o1"}, expectedErr: strconv.ErrSyntax},
		{name: "test_b", params: &models.StorageParams{Type: models.TypeMetricCounter,
			Name: "test", Value: "-1"}, expectedErr: models.ErrIncompatibleMetricValue},
	}
	for i, tt := range tests {
		t.Run(tt.name+"_"+strconv.Itoa(i), func(t *testing.T) {
			_storage := NewMemoryStorage()
			actualErr := _storage.MetricSet(tt.params)
			assert.ErrorIs(t, actualErr, tt.expectedErr)
		})
	}
}

func TestMemoryStorage_MetricSetWrongType(t *testing.T) {
	tests := []struct {
		expectedErr error
		params      *models.StorageParams
		name        string
	}{
		{name: "test_a", params: &models.StorageParams{Type: "",
			Name: "test", Value: int64(0)}, expectedErr: models.ErrNotImplementedMetricType},
		{name: "test_a", params: &models.StorageParams{Type: "empty",
			Name: "test", Value: int64(0)}, expectedErr: models.ErrNotImplementedMetricType},
	}
	for i, tt := range tests {
		t.Run(tt.name+"_"+strconv.Itoa(i), func(t *testing.T) {
			_storage := NewMemoryStorage()
			actualErr := _storage.MetricSet(tt.params)
			assert.ErrorIs(t, actualErr, tt.expectedErr)
		})
	}
}

func TestMemoryStorage_MetricSetEmptyName(t *testing.T) {
	tests := []struct {
		expectedErr error
		params      *models.StorageParams
		name        string
	}{
		{name: "test_a", params: &models.StorageParams{
			Type:  "",
			Name:  "",
			Value: int64(0),
		}, expectedErr: models.ErrEmptyMetricName},
	}
	for i, tt := range tests {
		t.Run(tt.name+"_"+strconv.Itoa(i), func(t *testing.T) {
			_storage := NewMemoryStorage()
			actualErr := _storage.MetricSet(tt.params)
			assert.ErrorIs(t, actualErr, tt.expectedErr)
		})
	}
}

func TestMemoryStorage_MetricGaugeSequentialSet(t *testing.T) {
	tests := []struct {
		params   *models.StorageParams
		name     string
		expected string
	}{
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: float64(0)}, expected: "0"},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: float64(3)}, expected: "3"},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: float64(6)}, expected: "6"},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: float64(9)}, expected: "9"},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: float64(0)}, expected: "0"},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: float64(-3)}, expected: "-3"},
	}

	_storage := NewMemoryStorage()

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

func TestMemoryStorage_MetricCounterSequentialSet(t *testing.T) {
	tests := []struct {
		params   *models.StorageParams
		name     string
		expected string
	}{
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricCounter,
			Name: "test", Value: int64(0)}, expected: "0"},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricCounter,
			Name: "test", Value: int64(3)}, expected: "3"},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricCounter,
			Name: "test", Value: int64(3)}, expected: "6"},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricCounter,
			Name: "test", Value: int64(3)}, expected: "9"},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricCounter,
			Name: "test", Value: int64(3)}, expected: "12"},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricCounter,
			Name: "test", Value: int64(3)}, expected: "15"},
	}

	_storage := NewMemoryStorage()

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

func TestMemoryStorage_MetricSequentialSet(t *testing.T) {
	tests := []struct {
		expectedErr error
		params      *models.StorageParams
		name        string
		expected    string
	}{
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: float64(0)}, expected: "0"},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricCounter,
			Name: "test", Value: int64(3)}, expected: "3"},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: float64(3)}, expected: "3"},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricCounter,
			Name: "test", Value: int64(3)}, expected: "6"},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test", Value: float64(-1)}, expected: "-1"},
		{name: "test_a", params: &models.StorageParams{Type: models.TypeMetricCounter,
			Name: "test", Value: int64(3)}, expected: "9"},

		{name: "test_b", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test1", Value: float64(888)}, expected: "888"},
		{name: "test_b", params: &models.StorageParams{Type: models.TypeMetricGauge,
			Name: "test1", Value: int64(-1)}, expected: "888", expectedErr: models.ErrIncompatibleMetricValueType},

		{name: "test_c", params: &models.StorageParams{Type: models.TypeMetricCounter,
			Name: "test2", Value: int64(999)}, expected: "999"},
		{name: "test_c", params: &models.StorageParams{Type: models.TypeMetricCounter,
			Name: "test2", Value: "-1"}, expected: "999", expectedErr: models.ErrIncompatibleMetricValue},
		{name: "test_c", params: &models.StorageParams{Type: models.TypeMetricCounter,
			Name: "test2", Value: int64(999)}, expected: "1998"},
		{name: "test_c", params: &models.StorageParams{Type: models.TypeMetricCounter,
			Name: "test2", Value: "-1"}, expected: "1998", expectedErr: models.ErrIncompatibleMetricValue},
	}

	_storage := NewMemoryStorage()

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
