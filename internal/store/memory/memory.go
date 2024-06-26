package memory

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/mcrgnt/yp1/internal/store/metric"
	"github.com/mcrgnt/yp1/internal/store/models"
)

type MemoryStorage struct {
	Metrics map[string]models.Metric
	emitter chan struct{}
	mu      sync.Mutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		Metrics: map[string]models.Metric{},
	}
}

func (t *MemoryStorage) isMetricExistsValue(params *models.StorageParams) (models.Metric, bool) {
	if v, ok := t.Metrics[params.Type+params.Name]; ok {
		return v, true
	}
	return nil, false
}

func (t *MemoryStorage) isMetricExists(params *models.StorageParams) (exists bool) {
	_, exists = t.isMetricExistsValue(params)
	return
}

func (t *MemoryStorage) metricSetNoLock(params *models.StorageParams) error {
	if params.Name == "" {
		return fmt.Errorf("metric set failed: %w", models.ErrEmptyMetricName)
	}
	if params.Type != models.TypeMetricCounter && params.Type != models.TypeMetricGauge {
		return fmt.Errorf("metric set failed: %w <%s>", models.ErrNotImplementedMetricType, params.Type)
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

func (t *MemoryStorage) MetricSet(params *models.StorageParams) error {
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

func (t *MemoryStorage) MetricReset(params *models.StorageParams) (err error) {
	t.mu.Lock()
	if t.isMetricExists(params) {
		t.Metrics[params.Type+params.Name].Reset()
	} else {
		err = fmt.Errorf("can't reset not existing metric: %s", params.Name)
	}
	t.mu.Unlock()
	return
}

func (t *MemoryStorage) GetMetricString(params *models.StorageParams) (err error) {
	t.mu.Lock()
	if v, ok := t.isMetricExistsValue(params); ok {
		params.String = v.String()
		params.Type = v.Type()
	} else {
		err = fmt.Errorf("get metric string: %w %s", models.ErrMetricNotFound, params.Name)
	}
	t.mu.Unlock()
	return
}

func (t *MemoryStorage) GetMetric(params *models.StorageParams) (err error) {
	t.mu.Lock()
	if v, ok := t.isMetricExistsValue(params); ok {
		params.Value = v.Value()
		params.Type = v.Type()
	} else {
		err = fmt.Errorf("get metric: %w %s", models.ErrMetricNotFound, params.Name)
	}
	t.mu.Unlock()
	return
}

func (t *MemoryStorage) GetMetricAll() (data string) {
	t.mu.Lock()
	for _, metric := range t.Metrics {
		data += metric.Name() + ": " + metric.String() + "\r\n"
	}
	t.mu.Unlock()
	return
}

func (t *MemoryStorage) SetAllJSON(data []byte) error {
	t.mu.Lock()
	defer func() {
		t.mu.Unlock()
	}()
	if err := json.Unmarshal(data, t); err != nil {
		return fmt.Errorf("unmarshal failed: %w", err)
	}
	return nil
}

func (t *MemoryStorage) GetAllJSON() ([]byte, error) {
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

func (t *MemoryStorage) UnmarshalJSON(data []byte) error {
	target := []*memStorageMarshaler{}
	if err := json.Unmarshal(data, &target); err != nil {
		return fmt.Errorf("unmarshal failed: %w", err)
	} else {
		for _, v := range target {
			if err := t.metricSetNoLock(&models.StorageParams{
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

func (t *MemoryStorage) MarshalJSON() ([]byte, error) {
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

func (t *MemoryStorage) Emitter() chan struct{} {
	t.emitter = make(chan struct{})
	return t.emitter
}
