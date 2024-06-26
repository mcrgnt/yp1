package metric

import (
	"fmt"

	"github.com/mcrgnt/yp1/internal/store/metric/counter"
	"github.com/mcrgnt/yp1/internal/store/metric/gauge"
	"github.com/mcrgnt/yp1/internal/store/models"
)

type NewMetricParams struct {
	Type  string
	Value any
	Name  string
}

func NewMetric(params *NewMetricParams) (models.Metric, error) {
	switch params.Type {
	case models.TypeMetricGauge:
		if m, err := gauge.NewGauge(&gauge.NewGaugeParams{
			Val:  params.Value,
			Name: params.Name,
		}); err != nil {
			return nil, fmt.Errorf("new gauge failed: %w", err)
		} else {
			return m, nil
		}
	case models.TypeMetricCounter:
		if m, err := counter.NewCounter(&counter.NewCounterParams{
			Val:  params.Value,
			Name: params.Name,
		}); err != nil {
			return nil, fmt.Errorf("new ccounter failed: %w", err)
		} else {
			return m, nil
		}
	default:
		return nil, fmt.Errorf("new metric: %w <%s>", models.ErrNotImplementedMetricType, params.Type)
	}
}
