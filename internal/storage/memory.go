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
	emitter chan struct{}
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

func (t *MemStorage) metricSetNoLock(params *StorageParams) error {
	if params.Name == "" {
		return fmt.Errorf("metric set failed: %w", common.ErrEmptyMetricName)
	}
	if params.Type != common.TypeMetricCounter && params.Type != common.TypeMetricGauge {
		return fmt.Errorf("metric set failed: %w <%s>", common.ErrNotImplementedMetricType, params.Type)
	}
	if t.isMetricExists(params) {
		if err := t.Metrics[params.Type+params.Name].Set(params.Value); err != nil {
			return fmt.Errorf("set failed: %w", err)
		}
	} else {
		if newMetric, err := metric.NewMetric(&metric.NewMetricParams{
			Type:  params.Type,
			Value: params.Value,
			Name:  params.Name,
		}); err != nil {
			return fmt.Errorf("new metric failed: %w", err)
		} else {
			t.Metrics[params.Type+params.Name] = newMetric
		}
	}
	return nil
}

func (t *MemStorage) MetricSet(params *StorageParams) error {
	t.mu.Lock()
	defer func() {
		t.mu.Unlock()
	}()
	if err := t.metricSetNoLock(params); err != nil {
		return fmt.Errorf("metric set failed: %w", err)
	}
	if t.emitter != nil {
		t.emitter <- struct{}{}
	}
	return nil
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
	for _, metric := range t.Metrics {
		data += metric.Name() + ": " + metric.String() + "\r\n"
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

type memStorageMarshaler struct {
	Value interface{} `json:"value"`
	Name  string      `json:"name"`
	Type  string      `json:"type"`
}

func (t *MemStorage) UnmarshalJSON(data []byte) error {
	target := []*memStorageMarshaler{}
	if err := json.Unmarshal(data, &target); err != nil {
		return fmt.Errorf("unmarshal failed: %w", err)
	} else {
		for _, v := range target {
			if err := t.metricSetNoLock(&StorageParams{
				Value: v.Value,
				Type:  v.Type,
				Name:  v.Name,
			}); err != nil {
				return fmt.Errorf("metric set failed: %w", err)
			}
		}
	}
	return nil
}

func (t *MemStorage) MarshalJSON() ([]byte, error) {
	target := []*memStorageMarshaler{}
	for _, v := range t.Metrics {
		target = append(target, &memStorageMarshaler{
			Name:  v.Name(),
			Type:  v.Type(),
			Value: v.Value(),
		})
	}
	if data, err := json.Marshal(&target); err != nil {
		return nil, fmt.Errorf("marshal failed: %w", err)
	} else {
		return data, nil
	}
}

func (t *MemStorage) Emitter() chan struct{} {
	t.emitter = make(chan struct{})
	return t.emitter
}
