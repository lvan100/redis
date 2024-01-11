package redis

import (
	"context"
	"fmt"
	"strconv"
	"unsafe"
)

type Cmdable interface {
	Driver() Driver
	CmdArgs() (ctx context.Context, cmd string, args []any)
}

type Replier[T Cmdable] struct{}

func (r *Replier[T]) Send() (any, error) {
	t := *(*T)(unsafe.Pointer(&r))
	ctx, cmd, args := t.CmdArgs()
	return t.Driver().Exec(ctx, cmd, args)
}

type StringReplier[T Cmdable] struct{ Replier[T] }

func (r *StringReplier[T]) Send() (string, error) {
	result, err := r.Replier.Send()
	if err != nil {
		return "", err
	}
	return result.(string), nil
}

type Result struct {
	Data []string
}

// NewResult returns a new *Result.
func NewResult(data ...string) *Result {
	return &Result{Data: data}
}

// Int executes a command whose reply is a `int64`.
func Int(v any, err error) (int64, error) {
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
		return Int(r.Data[0], nil)
	default:
		return 0, fmt.Errorf("redis: unexpected type (%T) for int64", v)
	}
}

// Float executes a command whose reply is a `float64`.
func Float(v any, err error) (float64, error) {
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
		return Float(r.Data[0], nil)
	default:
		return 0, fmt.Errorf("redis: unexpected type (%T) for float64", r)
	}
}

// String executes a command whose reply is a `string`.
func String(v any, err error) (string, error) {
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

// Slice executes a command whose reply is a `[]any`.
func Slice(v any, err error) ([]any, error) {
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
		return Slice(r.Data, nil)
	default:
		return nil, fmt.Errorf("redis: unexpected type (%T) for []any", v)
	}
}

// IntSlice executes a command whose reply is a `[]int64`.
func IntSlice(v any, err error) ([]int64, error) {
	slice, err := Slice(v, err)
	if err != nil {
		return nil, err
	}
	if len(slice) == 0 {
		return nil, nil
	}
	val := make([]int64, len(slice))
	for i, r := range slice {
		var n int64
		n, err = Int(r, nil)
		if err != nil {
			return nil, err
		}
		val[i] = n
	}
	return val, nil
}

// FloatSlice executes a command whose reply is a `[]float64`.
func FloatSlice(v any, err error) ([]float64, error) {
	slice, err := Slice(v, err)
	if err != nil {
		return nil, err
	}
	if len(slice) == 0 {
		return nil, nil
	}
	val := make([]float64, len(slice))
	for i, r := range slice {
		var f float64
		f, err = Float(r, nil)
		if err != nil {
			return nil, err
		}
		val[i] = f
	}
	return val, nil
}

// StringSlice executes a command whose reply is a `[]string`.
func StringSlice(v any, err error) ([]string, error) {
	slice, err := Slice(v, err)
	if err != nil {
		return nil, err
	}
	if len(slice) == 0 {
		return nil, nil
	}
	val := make([]string, len(slice))
	for i, r := range slice {
		var str string
		str, err = String(r, nil)
		if err != nil {
			return nil, err
		}
		val[i] = str
	}
	return val, nil
}

// StringMap executes a command whose reply is a `map[string]string`.
func StringMap(v any, err error) (map[string]string, error) {
	if err != nil {
		return nil, err
	}
	slice, err := StringSlice(v, err)
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

//// ZItemSlice executes a command whose reply is a `[]ZItem`.
//func ZItemSlice(v any, err error) ([]ZItem, error) {
//	if err != nil {
//		return nil, err
//	}
//	slice, err := StringSlice(v, err)
//	if err != nil {
//		return nil, err
//	}
//	if len(slice) == 0 {
//		return nil, nil
//	}
//	if len(slice)%2 != 0 {
//		return nil, fmt.Errorf("redis: unexpected slice length %d", len(slice))
//	}
//	val := make([]ZItem, len(slice)/2)
//	for i := 0; i < len(val); i++ {
//		idx := i * 2
//		var score float64
//		score, err = Float(slice[idx+1], nil)
//		if err != nil {
//			return nil, err
//		}
//		val[i].Member = slice[idx]
//		val[i].Score = score
//	}
//	return val, nil
//}
