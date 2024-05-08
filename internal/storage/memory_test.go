package storage

import (
	"reflect"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Equal_NewMemory(t *testing.T) {
	tests := []struct {
		name string
		want *Memory
	}{
		{
			name: "test0",
			want: &Memory{
				Gauges:   map[string]float64{},
				Counters: map[string]int64{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewMemory()
			assert.Equal(t, reflect.TypeOf(tt.want), reflect.TypeOf(actual))
			assert.Equal(t, tt.want.Gauges, actual.Gauges)
			assert.Equal(t, tt.want.Counters, actual.Counters)
		})
	}
}
func Test_NotEqual_NewMemory(t *testing.T) {
	tests := []struct {
		name string
		want *Memory
	}{
		{
			name: "test0",
			want: &Memory{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewMemory()
			assert.NotEqual(t, nil, actual)
			assert.NotEqual(t, tt.want.Gauges, actual.Gauges)
			assert.NotEqual(t, tt.want.Counters, actual.Counters)
		})
	}
}

func TestMemory_GaugesEqual_Update(t *testing.T) {
	type fields struct {
		Gauges   map[string]float64
		Counters map[string]int64
		//		mu       sync.Mutex
	}
	type args struct {
		params *StorageParams
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected *Memory
	}{
		{
			name: "test0",
			fields: fields{
				Gauges: map[string]float64{},
			},
			args: args{
				&StorageParams{
					Type:  "gauge",
					Name:  "test0",
					Value: float64(0),
				},
			},
			expected: &Memory{
				Gauges: map[string]float64{
					"test0": float64(0),
				},
			},
		},
		{
			name: "test-1",
			fields: fields{
				Gauges: map[string]float64{},
			},
			args: args{
				&StorageParams{
					Type:  "gauge",
					Name:  "test-1",
					Value: float64(-1),
				},
			},
			expected: &Memory{
				Gauges: map[string]float64{
					"test-1": float64(-1),
				},
			},
		},
		{
			name: "test1",
			fields: fields{
				Gauges: map[string]float64{},
			},
			args: args{
				&StorageParams{
					Type:  "gauge",
					Name:  "test1",
					Value: float64(1),
				},
			},
			expected: &Memory{
				Gauges: map[string]float64{
					"test1": float64(1),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Memory{
				Gauges:   tt.fields.Gauges,
				Counters: tt.fields.Counters,
				//				mu:       tt.fields.mu,
			}
			tr.Update(tt.args.params)
			assert.Equal(t, tt.fields.Gauges, tt.expected.Gauges)
		})
	}
}

func TestMemory_GaugesNotEqual_Update(t *testing.T) {
	type fields struct {
		Gauges   map[string]float64
		Counters map[string]int64
		//		mu       sync.Mutex
	}
	type args struct {
		params *StorageParams
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected *Memory
	}{
		{
			name: "test0",
			fields: fields{
				Gauges: map[string]float64{},
			},
			args: args{
				&StorageParams{
					Type:  "gauge",
					Name:  "test0",
					Value: float64(6.1),
				},
			},
			expected: &Memory{
				Gauges: map[string]float64{
					"test0": float64(6.0),
				},
			},
		},
		{
			name: "test1",
			fields: fields{
				Gauges: map[string]float64{},
			},
			args: args{
				&StorageParams{
					Type:  "gauge",
					Name:  "test1",
					Value: float64(6),
				},
			},
			expected: &Memory{
				Gauges: map[string]float64{
					"test1": float64(-6),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Memory{
				Gauges:   tt.fields.Gauges,
				Counters: tt.fields.Counters,
				//				mu:       tt.fields.mu,
			}
			tr.Update(tt.args.params)
			assert.NotEqual(t, tt.fields.Gauges, tt.expected.Gauges)
		})
	}
}

func TestMemory_CountersEqual_Update(t *testing.T) {
	type fields struct {
		Gauges   map[string]float64
		Counters map[string]int64
		//		mu       sync.Mutex
	}
	type args struct {
		params *StorageParams
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected *Memory
	}{
		{
			name: "test0",
			fields: fields{
				Counters: map[string]int64{},
			},
			args: args{
				&StorageParams{
					Type:  "counter",
					Name:  "test0",
					Value: int64(1),
				},
			},
			expected: &Memory{
				Counters: map[string]int64{
					"test0": int64(1),
				},
			},
		},
		{
			name: "test1",
			fields: fields{
				Counters: map[string]int64{},
			},
			args: args{
				&StorageParams{
					Type:  "counter",
					Name:  "test1",
					Value: int64(-6),
				},
			},
			expected: &Memory{
				Counters: map[string]int64{
					"test1": int64(-6),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Memory{
				Gauges:   tt.fields.Gauges,
				Counters: tt.fields.Counters,
				//				mu:       tt.fields.mu,
			}
			tr.Update(tt.args.params)
			assert.Equal(t, tt.fields.Counters, tt.expected.Counters)
		})
	}
}

func TestMemory_CountersNotEqual_Update(t *testing.T) {
	type fields struct {
		Gauges   map[string]float64
		Counters map[string]int64
		//		mu       sync.Mutex
	}
	type args struct {
		params *StorageParams
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected *Memory
	}{
		{
			name: "test0",
			fields: fields{
				Counters: map[string]int64{},
			},
			args: args{
				&StorageParams{
					Type:  "counter",
					Name:  "test0",
					Value: int64(-1),
				},
			},
			expected: &Memory{
				Counters: map[string]int64{
					"test0": int64(6),
				},
			},
		},
		{
			name: "test1",
			fields: fields{
				Counters: map[string]int64{},
			},
			args: args{
				&StorageParams{
					Type:  "counter",
					Name:  "test1",
					Value: int64(6),
				},
			},
			expected: &Memory{
				Counters: map[string]int64{
					"test1": int64(-6),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Memory{
				Gauges:   tt.fields.Gauges,
				Counters: tt.fields.Counters,
				//				mu:       tt.fields.mu,
			}
			tr.Update(tt.args.params)
			assert.NotEqual(t, tt.fields.Counters, tt.expected.Counters)
		})
	}
}

func TestMemory_Equal_Get(t *testing.T) {
	type fields struct {
		Gauges   map[string]float64
		Counters map[string]int64
		//		mu       sync.Mutex
	}
	type args struct {
		params *StorageParams
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected *StorageParams
	}{
		{
			name: "test-1",
			fields: fields{
				Gauges: map[string]float64{
					"test-1": -3.2,
				},
			},
			args: args{
				&StorageParams{
					Name: "test-1",
				},
			},
			expected: &StorageParams{
				Type:  "gauge",
				Name:  "test-1",
				Value: float64(-3.2),
			},
		},
		{
			name: "test0",
			fields: fields{
				Gauges: map[string]float64{
					"test0": 0,
				},
			},
			args: args{
				&StorageParams{
					Name: "test0",
				},
			},
			expected: &StorageParams{
				Type:  "gauge",
				Name:  "test0",
				Value: float64(0),
			},
		},
		{
			name: "test1",
			fields: fields{
				Gauges: map[string]float64{
					"test1": 3.2,
				},
			},
			args: args{
				&StorageParams{
					Name: "test1",
				},
			},
			expected: &StorageParams{
				Type:  "gauge",
				Name:  "test1",
				Value: float64(3.2),
			},
		},
		{
			name: "test2",
			fields: fields{
				Counters: map[string]int64{
					"test2": 3,
				},
			},
			args: args{
				&StorageParams{
					Name: "test2",
				},
			},
			expected: &StorageParams{
				Type:  "counter",
				Name:  "test2",
				Value: int64(3),
			},
		},
		{
			name: "test3",
			fields: fields{
				Counters: map[string]int64{
					"test3": 0,
				},
			},
			args: args{
				&StorageParams{
					Name: "test3",
				},
			},
			expected: &StorageParams{
				Type:  "counter",
				Name:  "test3",
				Value: int64(0),
			},
		},
		{
			name: "test4",
			fields: fields{
				Counters: map[string]int64{
					"test4": -1,
				},
			},
			args: args{
				&StorageParams{
					Name: "test4",
				},
			},
			expected: &StorageParams{
				Type:  "counter",
				Name:  "test4",
				Value: int64(-1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Memory{
				Gauges:   tt.fields.Gauges,
				Counters: tt.fields.Counters,
				//				mu:       tt.fields.mu,
			}
			tr.Get(tt.args.params)
			assert.Equal(t, *tt.args.params, *tt.expected)
		})
	}
}

func TestMemory_NotEqual_Get(t *testing.T) {
	type fields struct {
		Gauges   map[string]float64
		Counters map[string]int64
		mu       sync.Mutex
	}
	type args struct {
		params *StorageParams
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected *StorageParams
	}{
		{
			name: "test-1",
			fields: fields{
				Gauges: map[string]float64{
					"test-1": -3.2,
				},
			},
			args: args{
				&StorageParams{
					Name: "test-1",
				},
			},
			expected: &StorageParams{
				Type:  "counter",
				Name:  "test-1",
				Value: float64(-3.2),
			},
		},
		{
			name: "test0",
			fields: fields{
				Gauges: map[string]float64{
					"test0": 0,
				},
			},
			args: args{
				&StorageParams{
					Name: "test0",
				},
			},
			expected: &StorageParams{
				Type:  "gauge",
				Name:  "test0",
				Value: 0,
			},
		},
		{
			name: "test1",
			fields: fields{
				Gauges: map[string]float64{
					"test1": 3.2,
				},
			},
			args: args{
				&StorageParams{
					Name: "test1",
				},
			},
			expected: &StorageParams{
				Type:  "gauge",
				Name:  "test1~",
				Value: float64(3.2),
			},
		},
		{
			name: "test2",
			fields: fields{
				Counters: map[string]int64{
					"test2": 3,
				},
			},
			args: args{
				&StorageParams{
					Name: "test2",
				},
			},
			expected: &StorageParams{
				Type:  "gauge",
				Name:  "test2",
				Value: int64(3),
			},
		},
		{
			name: "test3",
			fields: fields{
				Counters: map[string]int64{
					"test3": 0,
				},
			},
			args: args{
				&StorageParams{
					Name: "test3",
				},
			},
			expected: &StorageParams{
				Type:  "counter",
				Name:  "",
				Value: int64(0),
			},
		},
		{
			name: "test4",
			fields: fields{
				Counters: map[string]int64{
					"test4": -1,
				},
			},
			args: args{
				&StorageParams{
					Name: "test4",
				},
			},
			expected: &StorageParams{
				Type:  "counter",
				Name:  "test4",
				Value: -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Memory{
				Gauges:   tt.fields.Gauges,
				Counters: tt.fields.Counters,
				mu:       tt.fields.mu,
			}
			tr.Get(tt.args.params)
			assert.NotEqual(t, *tt.args.params, *tt.expected)
		})
	}
}
