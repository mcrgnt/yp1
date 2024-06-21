package storage

import (
	"encoding/json"
	"fmt"

	"github.com/mcrgnt/yp1/internal/models"
)

type StorageParams struct {
	Value  any
	String string
	Type   string
	Name   string
}

type Storage interface {
	MetricSet(*StorageParams) error
	MetricReset(*StorageParams) error
	GetMetricString(*StorageParams) error
	GetMetric(*StorageParams) error
	GetMetricAll() string
	SetAllJSON([]byte) error
	GetAllJSON() ([]byte, error)
	Emitter() chan struct{}
}

type NewMemStorageParams struct {
	Type string
}

func NewStorage(params *NewMemStorageParams) (storage Storage) {
	switch params.Type {
	case "memory":
		storage = NewMemStorage()
	default:
		storage = NewMemStorage()
	}
	return
}

func (t *StorageParams) MarshalJSON() ([]byte, error) {
	temp := &models.Metrics{
		ID:    t.Name,
		MType: t.Type,
	}
	switch t.Type {
	case "counter":
		switch v := t.Value.(type) {
		case int64:
			temp.Delta = &v
		default:
			return nil, fmt.Errorf("wrong type: %T", v)
		}
		if bb, err := json.Marshal(temp); err != nil {
			return nil, fmt.Errorf("marshal: %w", err)
		} else {
			return bb, nil
		}
	case "gauge":
		switch v := t.Value.(type) {
		case float64:
			temp.Value = &v
		default:
			return nil, fmt.Errorf("wrong type: %T", v)
		}
		if bb, err := json.Marshal(temp); err != nil {
			return nil, fmt.Errorf("marshal: %w", err)
		} else {
			return bb, nil
		}
	default:
		return nil, fmt.Errorf("not implemented type: %s", t.Type)
	}
}

func (t *StorageParams) UnmarshalJSON(bb []byte) error {
	temp := &models.Metrics{}
	if err := json.Unmarshal(bb, temp); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	} else {
		t.Name = temp.ID
		t.Type = temp.MType
		switch t.Type {
		case "counter":
			if temp.Delta != nil {
				t.Value = *temp.Delta
			}
		case "gauge":
			if temp.Value != nil {
				t.Value = *temp.Value
			}
		default:
			return fmt.Errorf("not implemented type: %s", t.Type)
		}
	}
	return nil
}
