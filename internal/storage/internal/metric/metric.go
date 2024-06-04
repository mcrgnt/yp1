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
	Value() any
}

type NewMetricParams struct {
	Value any
	Type  string
}

func NewMetric(params *NewMetricParams) (Metric, error) {
	switch params.Type {
	case common.TypeMetricGauge:
		return NewGauge(&NewGaugeParams{
			Val: params.Value,
		})
	case common.TypeMetricCounter:
		return NewCounter(&NewCounterParams{
			Val: params.Value,
		})
	}
	return nil, fmt.Errorf("new metric: %w <%s>", common.ErrNotImplementedMetricType, params.Type)
}
