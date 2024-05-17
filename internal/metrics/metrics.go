package metrics

import (
	"fmt"
	"math/rand"
	"runtime"

	"github.com/mcrgnt/yp1/internal/common"
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
	{
		err := params.Storage.MetricSet(&storage.StorageParams{
			Type:  common.MetricTypeGauge,
			Name:  "RandomValue",
			Value: rand.Float64(),
		})
		if err != nil {
			fmt.Println(err)
		}
	}
	{
		err := params.Storage.MetricSet(&storage.StorageParams{
			Type:  common.MetricTypeCounter,
			Name:  "PollCount",
			Value: int64(1),
		})
		if err != nil {
			fmt.Println(err)
		}
	}
}

type ReportMetricsParams struct {
	Storage storage.Storage
	Address string
}

func ReportMetrics(params *ReportMetricsParams) {
	var err error
	for _type, names := range metricsTypeNames {
		for _, name := range names {
			storageParams := &storage.StorageParams{
				Type: _type,
				Name: name,
			}
			err = params.Storage.GetMetricString(storageParams)
			if err != nil {
				fmt.Println(err)
			}
			err = reporter.Report(&reporter.ReportParams{
				URL: fmt.Sprintf("http://%s/update/%s/%s/%v",
					params.Address,
					storageParams.Type,
					storageParams.Name,
					storageParams.String,
				),
			})
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	_ = params.Storage.MetricReset(&storage.StorageParams{
		Type: common.MetricTypeCounter,
		Name: "PollCount",
	})
}

var (
	metricsTypeNames = map[string][]string{}
)

func getFullMetricsGaugeNamesList() (metricsNamesList []string) {
	metricsNamesList = append(metricsNamesList, PollMetricsFromMemStatsList...)
	metricsNamesList = append(metricsNamesList, "RandomValue")
	return
}

func getFullMetricsCounterNamesList() (metricsNamesList []string) {
	metricsNamesList = append(metricsNamesList, "PollCount")
	return
}

func init() {
	metricsTypeNames[common.MetricTypeGauge] = getFullMetricsGaugeNamesList()
	metricsTypeNames[common.MetricTypeCounter] = getFullMetricsCounterNamesList()
}
