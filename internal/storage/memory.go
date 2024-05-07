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

func (t *Memory) Update(params *Update) {
	t.mu.Lock()
	switch params.Type {
	case "gauge":
		t.Gauges[params.Name] = params.Value.(float64)
	case "counter":
		value := t.Counters[params.Name]
		t.Counters[params.Name] = value + params.Value.(int64)
	}
	t.mu.Unlock()
}
