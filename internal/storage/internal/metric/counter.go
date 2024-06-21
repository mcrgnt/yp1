package metric

import (
	"fmt"
	"strconv"

	"github.com/mcrgnt/yp1/internal/common"
)

type Counter struct {
	name string
	val  int64
}

type NewCounterParams struct {
	Val  interface{}
	Name string
}

func fromAnyToInt64(value any) (int64, error) {
	switch v := value.(type) {
	case float64:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint64:
		return int64(v), nil
	case int64:
		return v, nil
	case *int64:
		return *v, nil
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

func fromAnyToInt64WithCheckForNegative(value any) (v int64, err error) {
	v, err = fromAnyToInt64(value)
	if err != nil {
		return
	}
	if v < 0 {
		err = fmt.Errorf("convert to int64: %w %T", common.ErrIncompatibleMetricValue, value)
	}
	return
}

func NewCounter(params *NewCounterParams) (counter *Counter, err error) {
	value, err := fromAnyToInt64WithCheckForNegative(params.Val)
	if err != nil {
		return nil, err
	}
	return &Counter{
		val:  value,
		name: params.Name,
	}, nil
}

func (t *Counter) Set(value any) (err error) {
	newValue, err := fromAnyToInt64WithCheckForNegative(value)
	if err != nil {
		return
	}
	t.val += newValue
	return
}

func (t *Counter) Reset() {
	t.val = 0
}

func (t *Counter) Type() string {
	return common.TypeMetricCounter
}

func (t *Counter) String() string {
	return strconv.FormatInt(t.val, 10)
}

func (t *Counter) Value() any {
	return t.val
}

func (t *Counter) Name() string {
	return t.name
}
