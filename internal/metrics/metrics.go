package metrics

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"runtime"

	"github.com/mcrgnt/yp1/internal/reporter"
	"github.com/mcrgnt/yp1/internal/store/models"
)

//go:generate go run ../../cmd/genPollMetrics/main.go

const (
	TypeMetricGauge   = "gauge"
	TypeMetricCounter = "counter"
)

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
	Storage models.Storage
}

func PollMetrics(params *PollMetricsParams) error {
	runtime.ReadMemStats(MemStats)
	if err := pollMetrics(params); err != nil {
		return fmt.Errorf("poll metrics failed: %w", err)
	}
	if err := params.Storage.MetricSet(&models.StorageParams{
		Type:  TypeMetricGauge,
		Name:  "RandomValue",
		Value: rand.Float64(),
	}); err != nil {
		return fmt.Errorf("metric set failed: %w", err)
	}
	if err := params.Storage.MetricSet(&models.StorageParams{
		Type:  TypeMetricCounter,
		Name:  "PollCount",
		Value: int64(1),
	}); err != nil {
		return fmt.Errorf("metric set failed: %w", err)
	}
	return nil
}

type ReportMetricsParams struct {
	Storage models.Storage
	Address string
}

func ReportMetrics(params *ReportMetricsParams) error {
	var err error
	for _type, names := range metricsTypeNames {
		for _, name := range names {
			storageParams := &models.StorageParams{
				Type: _type,
				Name: name,
			}
			if err = params.Storage.GetMetric(storageParams); err != nil {
				return fmt.Errorf("get metric failed: %w", err)
			}
			if data, err := json.Marshal(storageParams); err != nil {
				return fmt.Errorf("json marshal failed: %w", err)
			} else if err = reporter.Report(&reporter.ReportParams{
				URL: fmt.Sprintf("http://%s/update/",
					params.Address,
				),
				Body: data,
			}); err != nil {
				return fmt.Errorf("report failed: %w", err)
			}
		}
	}

	if err = params.Storage.MetricReset(&models.StorageParams{
		Type: TypeMetricCounter,
		Name: "PollCount",
	}); err != nil {
		return fmt.Errorf("metric reset failed: %w", err)
	}
	return nil
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
	metricsTypeNames[TypeMetricGauge] = getFullMetricsGaugeNamesList()
	metricsTypeNames[TypeMetricCounter] = getFullMetricsCounterNamesList()
}
