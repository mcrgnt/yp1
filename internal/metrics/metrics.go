package metrics

import (
	"fmt"
	"math/rand"
	"net/http"
	"reflect"
	"runtime"
	"sync/atomic"

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
	fmt.Println("poll")
	runtime.ReadMemStats(memStats)

	val := reflect.ValueOf(memStats).Elem()
	for _, name := range PollMetricsFromMemStatsList {
		updateParams := &storage.StorageParams{
			Type: "gauge",
			Name: name,
		}
		switch val.FieldByName(name).Interface().(type) {
		case uint32, uint64:
			updateParams.Value = float64(val.FieldByName(name).Uint())
		default:
			updateParams.Value = val.FieldByName(name).Float()

		}
		params.Storage.Update(updateParams)
	}
	{
		PollMetricsInc.Add(1)
		updateParams := &storage.StorageParams{
			Type:  "counter",
			Name:  "PollCount",
			Value: int64(PollMetricsInc.Load()),
		}
		params.Storage.Update(updateParams)
	}
	{
		updateParams := &storage.StorageParams{
			Type:  "gauge",
			Name:  "RandomValue",
			Value: rand.Float64(),
		}
		params.Storage.Update(updateParams)
	}
}

type ReportMetricsParams struct {
	Storage storage.MemStorage
	Address string
}

func ReportMetrics(params *ReportMetricsParams) (err error) {
	for _, name := range PollMetricsFromMemStatsList {
		storageParams := &storage.StorageParams{
			Name: name,
		}
		params.Storage.Get(storageParams)
		_, err = http.Post(fmt.Sprintf("http://%s/update/%s/%s/%v", params.Address, storageParams.Type, storageParams.Name, storageParams.Value), "text/plain", nil)
		if err != nil {
			err = fmt.Errorf("post to server: %v", err)
			return
		}
	}
	return
}
