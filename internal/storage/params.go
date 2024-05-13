package storage

import (
	"errors"
	"fmt"
	"strconv"
)

type StorageParams struct {
	Value any
	Type  string
	Name  string
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
		err = errors.New("validate name: empty metric name not allowed")
	}
	return
}

func (t *StorageParams) ValidateValue() (err error) {
	switch t.Type {
	case gauge:
		t.Value, err = strconv.ParseFloat(t.Value.(string), 64)
		if err != nil {
			err = fmt.Errorf("validate value: %w", err)
		}
	case counter:
		t.Value, err = strconv.ParseInt(t.Value.(string), 10, 64)
		if err != nil {
			err = fmt.Errorf("validate value: %w", err)
		}
	}
	return
}
