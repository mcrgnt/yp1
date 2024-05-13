package metrics

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"sync/atomic"

	"github.com/mcrgnt/yp1/internal/reporter"
	"github.com/mcrgnt/yp1/internal/storage"
)

var (
	memStats                    = &runtime.MemStats{}
	PollMetricsFromMemStatsList = []string{
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
	}
	PollMetricsInc atomic.Int64
)

type PollMetricsParams struct {
	Storage storage.MemStorage
}

func PollMetrics(params *PollMetricsParams) {
	runtime.ReadMemStats(memStats)
	val := reflect.ValueOf(memStats).Elem()
	for _, name := range PollMetricsFromMemStatsList {
		updateParams := &storage.StorageParams{
			Type: "gauge",
			Name: name,
		}
		switch val.FieldByName(name).Interface().(type) {
		case uint32, uint64:
			updateParams.ValueFloat64 = float64(val.FieldByName(name).Uint())
		default:
			updateParams.ValueFloat64 = val.FieldByName(name).Float()
		}
		params.Storage.Update(updateParams)
	}
	{
		PollMetricsInc.Add(1)
		updateParams := &storage.StorageParams{
			Type:       "counter",
			Name:       "PollCount",
			ValueInt64: PollMetricsInc.Load(),
		}
		params.Storage.Update(updateParams)
	}
	{
		updateParams := &storage.StorageParams{
			Type:         "gauge",
			Name:         "RandomValue",
			ValueFloat64: rand.Float64(),
		}
		params.Storage.Update(updateParams)
	}
}

type ReportMetricsParams struct {
	Storage storage.MemStorage
	Address string
}

func getFullMetricsNamesList() (metricsNamesList []string) {
	metricsNamesList = append(metricsNamesList, PollMetricsFromMemStatsList...)
	metricsNamesList = append(metricsNamesList, "PollCount", "RandomValue")
	return
}

func ReportMetrics(params *ReportMetricsParams) {
	for _, name := range getFullMetricsNamesList() {
		storageParams := &storage.StorageParams{
			Name: name,
		}
		params.Storage.GetByName(storageParams)
		reporter.Report(&reporter.ReportParams{
			URL: fmt.Sprintf("http://%s/update/%s/%s/%v",
				params.Address,
				storageParams.Type,
				storageParams.Name,
				storageParams.Value,
			),
		})
	}
}
