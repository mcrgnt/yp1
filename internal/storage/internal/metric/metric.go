package metric

import (
	"fmt"

	"github.com/mcrgnt/yp1/internal/common"
)

type Metric interface {
	Set(any) error
	Reset()
	Type() string
	String() string
}

type NewMetricParams struct {
	Value any
	Type  string
}

func NewMetric(params *NewMetricParams) (Metric, error) {
	fmt.Println("_0")
	switch params.Type {
	case common.MetricTypeGauge:
		fmt.Println("_1")
		return NewGauge(&NewGaugeParams{
			Value: params.Value,
		})
	case common.MetricTypeCounter:
		fmt.Println("_2")
		return NewCounter(&NewCounterParams{
			Value: params.Value,
		})
	}
	fmt.Println("_3")
	return nil, fmt.Errorf("new metric: %w <%s>", common.ErrNotImplementedMetricType, params.Type)
}
