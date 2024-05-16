package metrics

import (
	"fmt"
	"math/rand"
	"runtime"

	"github.com/mcrgnt/yp1/internal/reporter"
	"github.com/mcrgnt/yp1/internal/storage"
)

//go:generate go run ../../cmd/genPollMetrics/main.go

var (
	MemStats                    = &runtime.MemStats{}
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
)

type PollMetricsParams struct {
	Storage storage.Storage
}

func PollMetrics(params *PollMetricsParams) {
	runtime.ReadMemStats(MemStats)
	pollMetrics(params)
	_ = params.Storage.MetricSet(&storage.StorageParams{
		Type:  "gauge",
		Name:  "RandomValue",
		Value: rand.Float64(),
	})
	_ = params.Storage.MetricSet(&storage.StorageParams{
		Type:  "counter",
		Name:  "PollCount",
		Value: int64(1),
	})
}

type ReportMetricsParams struct {
	Storage storage.Storage
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
		_ = params.Storage.GetMetricStringByName(storageParams)
		reporter.Report(&reporter.ReportParams{
			URL: fmt.Sprintf("http://%s/update/%s/%s/%v",
				params.Address,
				storageParams.Type,
				storageParams.Name,
				storageParams.String,
			),
		})
	}
	_ = params.Storage.MetricReset(&storage.StorageParams{
		Name: "PollCount",
	})
}
