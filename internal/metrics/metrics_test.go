package metrics

import (
	"reflect"
	"testing"

	"github.com/mcrgnt/yp1/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestPollMetrics(t *testing.T) {
	type args struct {
		params *PollMetricsParams
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test0",
			args: args{
				params: &PollMetricsParams{
					Storage: storage.NewMemStorage(&storage.NewMemStorageParams{}),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PollMetrics(tt.args.params)
			assert.NotEqual(t, storage.Memory{}, tt.args.params.Storage)
		})
	}
}

func Test_getFullMetricsNamesList(t *testing.T) {
	tests := []struct {
		name                 string
		wantMetricsNamesList []string
	}{
		// TODO: Add test cases.
		{
			name: "test0",
			wantMetricsNamesList: []string{
				"Alloc",
				"BuckHashSys",
				"Frees",
				"GCCPUFraction",
				"GCSys",
				"HeapAlloc",
				"HeapIdle",
				"HeapInuse",
				"HeapObjects",
				"HeapReleased",
				"HeapSys",
				"LastGC",
				"Lookups",
				"MCacheInuse",
				"MCacheSys",
				"MSpanInuse",
				"MSpanSys",
				"Mallocs",
				"NextGC",
				"NumForcedGC",
				"NumGC",
				"OtherSys",
				"PauseTotalNs",
				"StackInuse",
				"StackSys",
				"Sys",
				"TotalAlloc",
				"PollCount",
				"RandomValue",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMetricsNamesList := getFullMetricsNamesList(); !reflect.DeepEqual(gotMetricsNamesList, tt.wantMetricsNamesList) {
				t.Errorf("getFullMetricsNamesList() = %v, want %v", gotMetricsNamesList, tt.wantMetricsNamesList)
			}
		})
	}
}
