package storage

import (
	"sync"
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
	case "gauge":
		t.Gauges[params.Name] = params.Value.(float64)
	case "counter":
		value := t.Counters[params.Name]
		t.Counters[params.Name] = value + params.Value.(int64)
	}
}

func (t *Memory) Get(params *StorageParams) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if v, ok := t.Gauges[params.Name]; ok {
		params.Type = "gauge"
		params.Value = v
		return
	}
	if v, ok := t.Counters[params.Name]; ok {
		params.Type = "counter"
		params.Value = v
		return
	}
}
