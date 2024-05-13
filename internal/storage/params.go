package storage

import (
	"errors"
	"fmt"
	"strconv"
)

type StorageParams struct {
	Value        any
	ValueString  string
	Type         string
	Name         string
	ValueFloat64 float64
	ValueInt64   int64
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
		t.ValueFloat64, err = strconv.ParseFloat(t.ValueString, 64)
		if err != nil {
			err = fmt.Errorf("validate value: %w", err)
		}
		t.Value = t.ValueFloat64
	case counter:
		t.ValueInt64, err = strconv.ParseInt(t.ValueString, 10, 64)
		if err != nil {
			err = fmt.Errorf("validate value: %w", err)
		}
		t.Value = t.ValueInt64
	}
	return
}
