package storage

import (
	"fmt"
	"sync"

	"github.com/mcrgnt/yp1/internal/common"
	"github.com/mcrgnt/yp1/internal/storage/internal/metric"
)

type MemStorage struct {
	Metrics map[string]metric.Metric
	mu      sync.Mutex
}

func NewMemory() *MemStorage {
	return &MemStorage{
		Metrics: map[string]metric.Metric{},
	}
}

func (t *MemStorage) isMetricExists(params *StorageParams) bool {
	if _, ok := t.Metrics[params.Name]; ok {
		return true
	}
	return false
}

func (t *MemStorage) MetricSet(params *StorageParams) (err error) {
	fmt.Println("0")
	if params.Name == "" {
		fmt.Println("1")
		return fmt.Errorf("new metric: %w", common.ErrEmptyMetricName)
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	if t.isMetricExists(params) {
		fmt.Println("2")
		err = t.Metrics[params.Name].Set(params.Value)
	} else {
		fmt.Println("3")
		var newMetric metric.Metric
		newMetric, err = metric.NewMetric(&metric.NewMetricParams{
			Type:  params.Type,
			Value: params.Value,
		})
		if err != nil {
			fmt.Println("4")
			return
		}
		t.Metrics[params.Name] = newMetric
		fmt.Println("5")
	}
	return
}

func (t *MemStorage) MetricReset(params *StorageParams) (err error) {
	t.mu.Lock()
	if t.isMetricExists(params) {
		t.Metrics[params.Name].Reset()
	} else {
		err = fmt.Errorf("can't reset not existing metric: %s", params.Name)
	}
	t.mu.Unlock()
	return
}

func (t *MemStorage) GetMetricStringByName(params *StorageParams) (err error) {
	t.mu.Lock()
	if v, ok := t.Metrics[params.Name]; ok {
		params.String = v.String()
		params.Type = v.Type()
	} else {
		err = fmt.Errorf("metric not found: %s", params.Name)
	}
	t.mu.Unlock()
	return
}

func (t *MemStorage) GetMetricAll() (data string) {
	t.mu.Lock()
	for name, metric := range t.Metrics {
		data += name + ": " + metric.String() + "\r\n"
	}
	t.mu.Unlock()
	return
}
