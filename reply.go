package redis

import (
	"context"
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
