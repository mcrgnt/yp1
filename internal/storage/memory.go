package storage

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/mcrgnt/yp1/internal/common"
	"github.com/mcrgnt/yp1/internal/storage/internal/metric"
)

type MemStorage struct {
	Metrics map[string]metric.Metric
	mu      sync.Mutex
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		Metrics: map[string]metric.Metric{},
	}
}

func (t *MemStorage) isMetricExistsValue(params *StorageParams) (metric.Metric, bool) {
	if v, ok := t.Metrics[params.Type+params.Name]; ok {
		return v, true
	}
	return nil, false
}

func (t *MemStorage) isMetricExists(params *StorageParams) (exists bool) {
	_, exists = t.isMetricExistsValue(params)
	return
}

func (t *MemStorage) MetricSet(params *StorageParams) (err error) {
	if params.Name == "" {
		return fmt.Errorf("metric set: %w", common.ErrEmptyMetricName)
	}

	if params.Type != common.TypeMetricCounter && params.Type != common.TypeMetricGauge {
		return fmt.Errorf("metric set: %w <%s>", common.ErrNotImplementedMetricType, params.Type)
	}

	t.mu.Lock()
	defer func() {
		t.mu.Unlock()
	}()

	if t.isMetricExists(params) {
		err = t.Metrics[params.Type+params.Name].Set(params.Value)
	} else {
		var newMetric metric.Metric
		newMetric, err = metric.NewMetric(&metric.NewMetricParams{
			Type:  params.Type,
			Value: params.Value,
		})
		if err != nil {
			return
		}
		t.Metrics[params.Type+params.Name] = newMetric
	}
	return
}

func (t *MemStorage) MetricReset(params *StorageParams) (err error) {
	t.mu.Lock()
	if t.isMetricExists(params) {
		t.Metrics[params.Type+params.Name].Reset()
	} else {
		err = fmt.Errorf("can't reset not existing metric: %s", params.Name)
	}
	t.mu.Unlock()
	return
}

func (t *MemStorage) GetMetricString(params *StorageParams) (err error) {
	t.mu.Lock()
	if v, ok := t.isMetricExistsValue(params); ok {
		params.String = v.String()
		params.Type = v.Type()
	} else {
		err = fmt.Errorf("get metric string: %w %s", common.ErrMetricNotFound, params.Name)
	}
	t.mu.Unlock()
	return
}

func (t *MemStorage) GetMetric(params *StorageParams) (err error) {
	t.mu.Lock()
	if v, ok := t.isMetricExistsValue(params); ok {
		params.Value = v.Value()
		params.Type = v.Type()
	} else {
		err = fmt.Errorf("get metric: %w %s", common.ErrMetricNotFound, params.Name)
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

func (t *MemStorage) SetAllJSON(data []byte) error {
	t.mu.Lock()
	defer func() {
		t.mu.Unlock()
	}()
	if err := json.Unmarshal(data, t); err != nil {
		return fmt.Errorf("unmarshal failed: %w", err)
	}
	return nil
}

func (t *MemStorage) GetAllJSON() ([]byte, error) {
	t.mu.Lock()
	defer func() {
		t.mu.Unlock()
	}()
	if data, err := json.Marshal(t); err != nil {
		return nil, fmt.Errorf("marshal failed: %w", err)
	} else {
		return data, nil
	}
}

type marshaler struct {
	Name  string
	Type  string
	Value interface{}
}

func (t *MemStorage) UnmarshalJSON(data []byte) error {
	fmt.Println(">>> UnmarshalJSON", t)
	return nil
}

func (t *MemStorage) MarshalJSON() ([]byte, error) {
	fmt.Println(">>> MarshalJSON", t)
	for k, v := range t.Metrics {
		fmt.Println(k, v.Type(), v.Value())
	}
	return []byte(`{}`), nil
}
