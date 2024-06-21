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
	Name() string
	Value() any
}

type NewMetricParams struct {
	Value any
	Type  string
	Name  string
}

func NewMetric(params *NewMetricParams) (Metric, error) {
	switch params.Type {
	case common.TypeMetricGauge:
		return NewGauge(&NewGaugeParams{
			Val:  params.Value,
			Name: params.Name,
		})
	case common.TypeMetricCounter:
		return NewCounter(&NewCounterParams{
			Val:  params.Value,
			Name: params.Name,
		})
	}
	return nil, fmt.Errorf("new metric: %w <%s>", common.ErrNotImplementedMetricType, params.Type)
}
