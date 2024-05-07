package storage

import (
	"fmt"
	"strconv"
)

type StorageParams struct {
	Type  string
	Name  string
	Value any
}

func (t *StorageParams) ValidateType() (err error) {
	switch t.Type {
	case "gauge", "counter":
	default:
		err = fmt.Errorf("validate type: unknown type: %s", t.Type)
	}
	return
}

func (t *StorageParams) ValidateName() (err error) {
	if t.Name == "" {
		err = fmt.Errorf("validate name: empty metric name not allowed")
	}
	return
}

func (t *StorageParams) ValidateValue() (err error) {
	switch t.Type {
	case "gauge":
		t.Value, err = strconv.ParseFloat(t.Value.(string), 64)
	case "counter":
		t.Value, err = strconv.ParseInt(t.Value.(string), 10, 64)
	}
	if err != nil {
		err = fmt.Errorf("validate value: %v", err)
	}
	return
}
