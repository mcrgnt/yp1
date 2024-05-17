package common

import "errors"

const (
	TypeMetricGauge   = "gauge"
	TypeMetricCounter = "counter"
)

var (
	ErrNotImplementedMetricType    = errors.New("not implemented metric type")
	ErrIncompatibleMetricValueType = errors.New("incompatible metric value type")
	ErrIncompatibleMetricValue     = errors.New("incompatible metric value")
	ErrEmptyMetricName             = errors.New("empty metric name")
	ErrMetricNotFound              = errors.New("metric not found")
)
