package storage

import (
	"fmt"
	"strconv"
	"sync"
)

const (
	gauge   = "gauge"
	counter = "counter"
)

type Memory struct {
	Gauges   map[string]float64
	Counters map[string]int64
	mu       sync.Mutex
}

func NewMemory() *Memory {
	return &Memory{
		Gauges:   map[string]float64{},
		Counters: map[string]int64{},
	}
}

func (t *Memory) Update(params *StorageParams) {
	t.mu.Lock()
	defer t.mu.Unlock()
	switch params.Type {
	case gauge:
		t.Gauges[params.Name] = params.Value.(float64) //nolint:all // type is checked in validate or pre update function
	case counter:
		t.Counters[params.Name] += params.Value.(int64) //nolint:all // type is checked in validate or pre update function
	}
}

func (t *Memory) Reset(params *StorageParams) {
	t.mu.Lock()
	defer t.mu.Unlock()
	switch params.Type {
	case gauge:
		t.Gauges[params.Name] = 0
	case counter:
		t.Counters[params.Name] = 0
	}
}

func (t *Memory) GetByName(params *StorageParams) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if v, ok := t.Gauges[params.Name]; ok {
		params.Type = gauge
		params.Value = v
		return
	}
	if v, ok := t.Counters[params.Name]; ok {
		params.Type = counter
		params.Value = v
		return
	}
}

func (t *Memory) GetByType(params *StorageParams) (value string, err error) {
	t.mu.Lock()
	defer t.mu.Unlock()
	switch params.Type {
	case gauge:
		if v, ok := t.Gauges[params.Name]; ok {
			params.Value = v
			value = strconv.FormatFloat(v, 'f', -1, 64)
			return
		}
	case counter:
		if v, ok := t.Counters[params.Name]; ok {
			params.Value = v
			value = strconv.FormatInt(v, 10)
			return
		}
	}
	err = fmt.Errorf("metric not found: %s", params.Name)
	return
}

func (t *Memory) GetAll() (data string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	for k, v := range t.Gauges {
		data += fmt.Sprintf("%s: %v\r\n", k, v)
	}
	for k, v := range t.Counters {
		data += fmt.Sprintf("%s: %v\r\n", k, v)
	}
	return
}
