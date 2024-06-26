package models

const (
	TypeMetricGauge   = "gauge"
	TypeMetricCounter = "counter"
)

type Metric interface {
	Set(value any) error
	Reset()
	Type() string
	String() string
	Value() any
	Name() string
}
