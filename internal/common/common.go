package common

import "errors"

const (
	MetricTypeGauge   = "gauge"
	MetricTypeCounter = "counter"
)

var (
	ErrNotImplementedMetricType    = errors.New("not implemented metric type")
	ErrIncompatibleMetricValueType = errors.New("incompatible metric value type")
	ErrEmptyMetricName             = errors.New("empty metric name")
	ErrMetricNotFound              = errors.New("metric not found")
)
