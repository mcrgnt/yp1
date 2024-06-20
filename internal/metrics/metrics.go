package metrics

import (
	"encoding/json"
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
			Type:  common.TypeMetricGauge,
			Name:  "RandomValue",
			Value: rand.Float64(),
		})
		if err != nil {
			fmt.Println(err)
		}
	}
	{
		err := params.Storage.MetricSet(&storage.StorageParams{
			Type:  common.TypeMetricCounter,
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
			err = params.Storage.GetMetric(storageParams)
			if err != nil {
				fmt.Println(err)
				return
			}
			if data, err := json.Marshal(storageParams); err != nil {
				fmt.Println(err)
				return
			} else {
				err = reporter.Report(&reporter.ReportParams{
					URL: fmt.Sprintf("http://%s/update/",
						params.Address,
					),
					Body: data,
				})
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}

	err = params.Storage.MetricReset(&storage.StorageParams{
		Type: common.TypeMetricCounter,
		Name: "PollCount",
	})
	if err != nil {
		fmt.Println(err)
	}
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
	metricsTypeNames[common.TypeMetricGauge] = getFullMetricsGaugeNamesList()
	metricsTypeNames[common.TypeMetricCounter] = getFullMetricsCounterNamesList()
}
