package storage

import (
	"fmt"
	"sync"

	"github.com/mcrgnt/yp1/internal/storage/internal/metric"
)

type Memory struct {
	Metrics map[string]metric.Metric
	mu      sync.Mutex
}

func NewMemory() *Memory {
	return &Memory{
		Metrics: map[string]metric.Metric{},
	}
}

func (t *Memory) isMetricExists(params *StorageParams) bool {
	if _, ok := t.Metrics[params.Name]; ok {
		return true
	}
	return false
}

func (t *Memory) MetricSet(params *StorageParams) (err error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.isMetricExists(params) {
		err = t.Metrics[params.Name].Set(params.Value)
	} else {
		var newMetric metric.Metric
		newMetric, err = metric.NewMetric(&metric.NewMetricParams{
			Type:  params.Type,
			Value: params.Value,
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		t.Metrics[params.Name] = newMetric
	}
	return
}

func (t *Memory) MetricReset(params *StorageParams) (err error) {
	t.mu.Lock()
	if t.isMetricExists(params) {
		t.Metrics[params.Name].Reset()
	} else {
		err = fmt.Errorf("can't reset not existing metric: %s", params.Name)
	}
	t.mu.Unlock()
	return
}

func (t *Memory) GetMetricStringByName(params *StorageParams) (err error) {
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

func (t *Memory) GetMetricAll() (data string) {
	t.mu.Lock()
	for name, metric := range t.Metrics {
		data += name + ": " + metric.String() + "\r\n"
	}
	t.mu.Unlock()
	return
}
