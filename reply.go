package redis

import (
	"fmt"
	"strconv"
)

type Int64Replier struct{ command }

func (r *Int64Replier) Result() (int64, error) {
	return toInt64(r.driver.Exec(r.ctx, r.cmd, r.args))
}

type Float64Replier struct{ command }

func (r *Float64Replier) Result() (float64, error) {
	return toFloat64(r.driver.Exec(r.ctx, r.cmd, r.args))
}

type StringReplier struct{ command }

func (r *StringReplier) Result() (string, error) {
	return toString(r.driver.Exec(r.ctx, r.cmd, r.args))
}

type SliceReplier struct{ command }

func (r *SliceReplier) Result() ([]any, error) {
	return toSlice(r.driver.Exec(r.ctx, r.cmd, r.args))
}

type Int64SliceReplier struct{ command }

func (r *Int64SliceReplier) Result() ([]int64, error) {
	return toInt64Slice(r.driver.Exec(r.ctx, r.cmd, r.args))
}

type Float64SliceReplier struct{ command }

func (r *Float64SliceReplier) Result() ([]float64, error) {
	return toFloat64Slice(r.driver.Exec(r.ctx, r.cmd, r.args))
}

type StringSliceReplier struct{ command }

func (r *StringSliceReplier) Result() ([]string, error) {
	return toStringSlice(r.driver.Exec(r.ctx, r.cmd, r.args))
}

type StringMapReplier struct{ command }

func (r *StringMapReplier) Result() (map[string]string, error) {
	return toStringMap(r.driver.Exec(r.ctx, r.cmd, r.args))
}

type ZItemSliceReplier struct{ command }

func (r *ZItemSliceReplier) Result() ([]ZItem, error) {
	return toZItemSlice(r.driver.Exec(r.ctx, r.cmd, r.args))
}

type Result struct {
	Data []string
}

// NewResult returns a new *Result.
func NewResult(data ...string) *Result {
	return &Result{Data: data}
}

// toInt64 executes a command whose reply is a `int64`.
func toInt64(v any, err error) (int64, error) {
	if err != nil {
		return 0, err
	}
	switch r := v.(type) {
	case nil:
		return 0, nil
	case int64:
		return r, nil
	case float64:
		return int64(r), nil
	case string:
		return strconv.ParseInt(r, 10, 64)
	case *Result:
		if len(r.Data) == 0 {
			return 0, fmt.Errorf("redis: no data")
		}
		return toInt64(r.Data[0], nil)
	default:
		return 0, fmt.Errorf("redis: unexpected type (%T) for int64", v)
	}
}

// toFloat64 executes a command whose reply is a `float64`.
func toFloat64(v any, err error) (float64, error) {
	if err != nil {
		return 0, err
	}
	switch r := v.(type) {
	case nil:
		return 0, nil
	case float64:
		return r, nil
	case int64:
		return float64(r), nil
	case string:
		return strconv.ParseFloat(r, 64)
	case *Result:
		if len(r.Data) == 0 {
			return 0, fmt.Errorf("redis: no data")
		}
		return toFloat64(r.Data[0], nil)
	default:
		return 0, fmt.Errorf("redis: unexpected type (%T) for float64", r)
	}
}

// toString executes a command whose reply is a `string`.
func toString(v any, err error) (string, error) {
	if err != nil {
		return "", err
	}
	switch r := v.(type) {
	case nil:
		return "", nil
	case string:
		return r, nil
	case *Result:
		if len(r.Data) == 0 {
			return "", fmt.Errorf("redis: no data")
		}
		return r.Data[0], nil
	default:
		return "", fmt.Errorf("redis: unexpected type (%T) for string", v)
	}
}

// toSlice executes a command whose reply is a `[]any`.
func toSlice(v any, err error) ([]any, error) {
	if err != nil {
		return nil, err
	}
	switch r := v.(type) {
	case nil:
		return nil, nil
	case []any:
		return r, nil
	case []string:
		if len(r) == 0 {
			return nil, nil
		}
		slice := make([]any, len(r))
		for i, str := range r {
			if str == "NULL" {
				slice[i] = nil
			} else {
				slice[i] = str
			}
		}
		return slice, nil
	case *Result:
		return toSlice(r.Data, nil)
	default:
		return nil, fmt.Errorf("redis: unexpected type (%T) for []any", v)
	}
}

// toInt64Slice executes a command whose reply is a `[]int64`.
func toInt64Slice(v any, err error) ([]int64, error) {
	slice, err := toSlice(v, err)
	if err != nil {
		return nil, err
	}
	if len(slice) == 0 {
		return nil, nil
	}
	val := make([]int64, len(slice))
	for i, r := range slice {
		var n int64
		n, err = toInt64(r, nil)
		if err != nil {
			return nil, err
		}
		val[i] = n
	}
	return val, nil
}

// toFloat64Slice executes a command whose reply is a `[]float64`.
func toFloat64Slice(v any, err error) ([]float64, error) {
	slice, err := toSlice(v, err)
	if err != nil {
		return nil, err
	}
	if len(slice) == 0 {
		return nil, nil
	}
	val := make([]float64, len(slice))
	for i, r := range slice {
		var f float64
		f, err = toFloat64(r, nil)
		if err != nil {
			return nil, err
		}
		val[i] = f
	}
	return val, nil
}

// toStringSlice executes a command whose reply is a `[]string`.
func toStringSlice(v any, err error) ([]string, error) {
	slice, err := toSlice(v, err)
	if err != nil {
		return nil, err
	}
	if len(slice) == 0 {
		return nil, nil
	}
	val := make([]string, len(slice))
	for i, r := range slice {
		var str string
		str, err = toString(r, nil)
		if err != nil {
			return nil, err
		}
		val[i] = str
	}
	return val, nil
}

// toStringMap executes a command whose reply is a `map[string]string`.
func toStringMap(v any, err error) (map[string]string, error) {
	if err != nil {
		return nil, err
	}
	slice, err := toStringSlice(v, err)
	if err != nil {
		return nil, err
	}
	if len(slice) == 0 {
		return nil, nil
	}
	if len(slice)%2 != 0 {
		return nil, fmt.Errorf("redis: unexpected slice length %d", len(slice))
	}
	val := make(map[string]string, len(slice)/2)
	for i := 0; i < len(slice); i += 2 {
		val[slice[i]] = slice[i+1]
	}
	return val, nil
}

// toZItemSlice executes a command whose reply is a `[]ZItem`.
func toZItemSlice(v any, err error) ([]ZItem, error) {
	if err != nil {
		return nil, err
	}
	slice, err := toStringSlice(v, err)
	if err != nil {
		return nil, err
	}
	if len(slice) == 0 {
		return nil, nil
	}
	if len(slice)%2 != 0 {
		return nil, fmt.Errorf("redis: unexpected slice length %d", len(slice))
	}
	val := make([]ZItem, len(slice)/2)
	for i := 0; i < len(val); i++ {
		idx := i * 2
		var score float64
		score, err = toFloat64(slice[idx+1], nil)
		if err != nil {
			return nil, err
		}
		val[i].Member = slice[idx]
		val[i].Score = score
	}
	return val, nil
}
