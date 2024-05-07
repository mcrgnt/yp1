package storage

import (
	"fmt"
	"strconv"
)

type Update struct {
	Type  string
	Name  string
	Value any
}

func (t *Update) ValidateType() (err error) {
	switch t.Type {
	case "gauge", "counter":
	default:
		err = fmt.Errorf("validate type: unknown type: %s", t.Type)
	}
	return
}

func (t *Update) ValidateName() (err error) {
	if t.Name == "" {
		err = fmt.Errorf("validate name: empty metric name not allowed")
	}
	return
}

func (t *Update) ValidateValue() (err error) {
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
