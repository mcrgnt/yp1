package metric

import (
	"fmt"
	"strconv"

	"github.com/mcrgnt/yp1/internal/common"
)

type Counter struct {
	Value int64
}

type NewCounterParams struct {
	Value interface{}
}

func fromAnyToInt64(value any) (int64, error) {
	switch v := value.(type) {
	case uint32:
		return int64(v), nil
	case uint64:
		return int64(v), nil
	case int64:
		return v, nil
	case string:
		_v, e := strconv.ParseInt(v, 10, 64)
		if e != nil {
			return 0, fmt.Errorf("convert to int64: %w", e)
		}
		return _v, nil
	default:
		return 0, fmt.Errorf("convert to int64: %w %T", common.ErrIncompatibleMetricValueType, value)
	}
}

func NewCounter(params *NewCounterParams) (counter *Counter, err error) {
	value, err := fromAnyToInt64(params.Value)
	if err != nil {
		return nil, err
	}
	return &Counter{
		Value: value,
	}, nil
}

func (t *Counter) Set(value any) (err error) {
	t.Value, err = fromAnyToInt64(value)
	return
}

func (t *Counter) Reset() {
	t.Value = 0
}

func (t *Counter) Type() string {
	return common.MetricTypeCounter
}

func (t *Counter) String() string {
	return strconv.FormatInt(t.Value, 10)
}
