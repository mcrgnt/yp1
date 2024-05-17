package metric

import (
	"fmt"
	"strconv"

	"github.com/mcrgnt/yp1/internal/common"
)

type Gauge struct {
	Value float64
}

type NewGaugeParams struct {
	Value interface{}
}

func fromAnyToFloat64(value any) (float64, error) {
	switch v := value.(type) {
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case float64:
		return v, nil
	case string:
		_v, e := strconv.ParseFloat(v, 64)
		fmt.Println(":::::::::::::::::::::::::: PARSE:", v, "->", _v)
		if e != nil {
			return 0, fmt.Errorf("convert to float64: %w", e)
		}
		return _v, nil
	default:
		return 0, fmt.Errorf("convert to float64: %w %T", common.ErrIncompatibleMetricValueType, value)
	}
}

func NewGauge(params *NewGaugeParams) (*Gauge, error) {
	value, err := fromAnyToFloat64(params.Value)
	if err != nil {
		return nil, err
	}
	return &Gauge{
		Value: value,
	}, nil
}

func (t *Gauge) Set(value any) (err error) {
	t.Value, err = fromAnyToFloat64(value)
	return
}

func (t *Gauge) Reset() {
	t.Value = 0
}

func (t *Gauge) Type() string {
	return common.MetricTypeGauge
}

func (t *Gauge) String() string {
	fmt.Println("FORMAT:", t.Value, strconv.FormatFloat(t.Value, 'f', -1, 64))
	return strconv.FormatFloat(t.Value, 'f', -1, 64)
}
