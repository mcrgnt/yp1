// Code generated by go generate; DO NOT EDIT.
package metrics

import (
	"fmt"

	"github.com/mcrgnt/yp1/internal/storage"
	"github.com/mcrgnt/yp1/internal/common"
)

func pollMetrics(params *PollMetricsParams) {
	{
		err := params.Storage.MetricSet(&storage.StorageParams{
			Type: common.MetricTypeGauge,
			Name: "Alloc",
			Value: float64(MemStats.Alloc),
		})
		if err !=nil {
			fmt.Println(err)
		}
	}
	{
		err := params.Storage.MetricSet(&storage.StorageParams{
			Type: common.MetricTypeGauge,
			Name: "BuckHashSys",
			Value: float64(MemStats.BuckHashSys),
		})
		if err !=nil {
			fmt.Println(err)
		}
	}
	{
		err := params.Storage.MetricSet(&storage.StorageParams{
			Type: common.MetricTypeGauge,
			Name: "Frees",
			Value: float64(MemStats.Frees),
		})
		if err !=nil {
			fmt.Println(err)
		}
	}
	{
		err := params.Storage.MetricSet(&storage.StorageParams{
			Type: common.MetricTypeGauge,
			Name: "GCCPUFraction",
			Value: MemStats.GCCPUFraction,
		})
		if err !=nil {
			fmt.Println(err)
		}
	}
	{
		err := params.Storage.MetricSet(&storage.StorageParams{
			Type: common.MetricTypeGauge,
			Name: "GCSys",
			Value: float64(MemStats.GCSys),
		})
		if err !=nil {
			fmt.Println(err)
		}
	}
	{
		err := params.Storage.MetricSet(&storage.StorageParams{
			Type: common.MetricTypeGauge,
			Name: "HeapAlloc",
			Value: float64(MemStats.HeapAlloc),
		})
		if err !=nil {
			fmt.Println(err)
		}
	}
	{
		err := params.Storage.MetricSet(&storage.StorageParams{
			Type: common.MetricTypeGauge,
			Name: "HeapIdle",
			Value: float64(MemStats.HeapIdle),
		})
		if err !=nil {
			fmt.Println(err)
		}
	}
	{
		err := params.Storage.MetricSet(&storage.StorageParams{
			Type: common.MetricTypeGauge,
			Name: "HeapInuse",
			Value: float64(MemStats.HeapInuse),
		})
		if err !=nil {
			fmt.Println(err)
		}
	}
	{
		err := params.Storage.MetricSet(&storage.StorageParams{
			Type: common.MetricTypeGauge,
			Name: "HeapObjects",
			Value: float64(MemStats.HeapObjects),
		})
		if err !=nil {
			fmt.Println(err)
		}
	}
	{
		err := params.Storage.MetricSet(&storage.StorageParams{
			Type: common.MetricTypeGauge,
			Name: "HeapReleased",
			Value: float64(MemStats.HeapReleased),
		})
		if err !=nil {
			fmt.Println(err)
		}
	}
	{
		err := params.Storage.MetricSet(&storage.StorageParams{
			Type: common.MetricTypeGauge,
			Name: "HeapSys",
			Value: float64(MemStats.HeapSys),
		})
		if err !=nil {
			fmt.Println(err)
		}
	}
	{
		err := params.Storage.MetricSet(&storage.StorageParams{
			Type: common.MetricTypeGauge,
			Name: "LastGC",
			Value: float64(MemStats.LastGC),
		})
		if err !=nil {
			fmt.Println(err)
		}
	}
	{
		err := params.Storage.MetricSet(&storage.StorageParams{
			Type: common.MetricTypeGauge,
			Name: "Lookups",
			Value: float64(MemStats.Lookups),
		})
		if err !=nil {
			fmt.Println(err)
		}
	}
	{
		err := params.Storage.MetricSet(&storage.StorageParams{
			Type: common.MetricTypeGauge,
			Name: "MCacheInuse",
			Value: float64(MemStats.MCacheInuse),
		})
		if err !=nil {
			fmt.Println(err)
		}
	}
	{
		err := params.Storage.MetricSet(&storage.StorageParams{
			Type: common.MetricTypeGauge,
			Name: "MCacheSys",
			Value: float64(MemStats.MCacheSys),
		})
		if err !=nil {
			fmt.Println(err)
		}
	}
	{
		err := params.Storage.MetricSet(&storage.StorageParams{
			Type: common.MetricTypeGauge,
			Name: "MSpanInuse",
			Value: float64(MemStats.MSpanInuse),
		})
		if err !=nil {
			fmt.Println(err)
		}
	}
	{
		err := params.Storage.MetricSet(&storage.StorageParams{
			Type: common.MetricTypeGauge,
			Name: "MSpanSys",
			Value: float64(MemStats.MSpanSys),
		})
		if err !=nil {
			fmt.Println(err)
		}
	}
	{
		err := params.Storage.MetricSet(&storage.StorageParams{
			Type: common.MetricTypeGauge,
			Name: "Mallocs",
			Value: float64(MemStats.Mallocs),
		})
		if err !=nil {
			fmt.Println(err)
		}
	}
	{
		err := params.Storage.MetricSet(&storage.StorageParams{
			Type: common.MetricTypeGauge,
			Name: "NextGC",
			Value: float64(MemStats.NextGC),
		})
		if err !=nil {
			fmt.Println(err)
		}
	}
	{
		err := params.Storage.MetricSet(&storage.StorageParams{
			Type: common.MetricTypeGauge,
			Name: "NumForcedGC",
			Value: float64(MemStats.NumForcedGC),
		})
		if err !=nil {
			fmt.Println(err)
		}
	}
	{
		err := params.Storage.MetricSet(&storage.StorageParams{
			Type: common.MetricTypeGauge,
			Name: "NumGC",
			Value: float64(MemStats.NumGC),
		})
		if err !=nil {
			fmt.Println(err)
		}
	}
	{
		err := params.Storage.MetricSet(&storage.StorageParams{
			Type: common.MetricTypeGauge,
			Name: "OtherSys",
			Value: float64(MemStats.OtherSys),
		})
		if err !=nil {
			fmt.Println(err)
		}
	}
	{
		err := params.Storage.MetricSet(&storage.StorageParams{
			Type: common.MetricTypeGauge,
			Name: "PauseTotalNs",
			Value: float64(MemStats.PauseTotalNs),
		})
		if err !=nil {
			fmt.Println(err)
		}
	}
	{
		err := params.Storage.MetricSet(&storage.StorageParams{
			Type: common.MetricTypeGauge,
			Name: "StackInuse",
			Value: float64(MemStats.StackInuse),
		})
		if err !=nil {
			fmt.Println(err)
		}
	}
	{
		err := params.Storage.MetricSet(&storage.StorageParams{
			Type: common.MetricTypeGauge,
			Name: "StackSys",
			Value: float64(MemStats.StackSys),
		})
		if err !=nil {
			fmt.Println(err)
		}
	}
	{
		err := params.Storage.MetricSet(&storage.StorageParams{
			Type: common.MetricTypeGauge,
			Name: "Sys",
			Value: float64(MemStats.Sys),
		})
		if err !=nil {
			fmt.Println(err)
		}
	}
	{
		err := params.Storage.MetricSet(&storage.StorageParams{
			Type: common.MetricTypeGauge,
			Name: "TotalAlloc",
			Value: float64(MemStats.TotalAlloc),
		})
		if err !=nil {
			fmt.Println(err)
		}
	}
}
