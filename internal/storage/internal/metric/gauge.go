package metric

import (
	"fmt"
	"strconv"

	"github.com/mcrgnt/yp1/internal/common"
)

type Gauge struct {
	Val float64
}

type NewGaugeParams struct {
	Val interface{}
}

func fromAnyToFloat64(value any) (float64, error) {
	switch v := value.(type) {
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case float64:
		return v, nil
	case *float64:
		return *v, nil
	case string:
		_v, e := strconv.ParseFloat(v, 64)
		if e != nil {
			return 0, fmt.Errorf("convert to float64: %w", e)
		}
		return _v, nil
	default:
		return 0, fmt.Errorf("convert to float64: %w %T", common.ErrIncompatibleMetricValueType, value)
	}
}

func NewGauge(params *NewGaugeParams) (*Gauge, error) {
	value, err := fromAnyToFloat64(params.Val)
	if err != nil {
		return nil, err
	}
	return &Gauge{
		Val: value,
	}, nil
}

func (t *Gauge) Set(value any) (err error) {
	v, err := fromAnyToFloat64(value)
	if err != nil {
		return
	}
	t.Val = v
	return
}

func (t *Gauge) Reset() {
	t.Val = 0
}

func (t *Gauge) Type() string {
	return common.TypeMetricGauge
}

func (t *Gauge) String() string {
	return strconv.FormatFloat(t.Val, 'f', -1, 64)
}

func (t *Gauge) Value() any {
	return t.Val
}
