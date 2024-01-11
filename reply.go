package redis

import (
	"context"
	"fmt"
	"unsafe"
)

type Cmdable interface {
	Cmd() (ctx context.Context, cmd string, args []any)
}

type StringReplier[T Cmdable] struct{}

func (r *StringReplier[T]) Send() (string, error) {
	t := *(*T)(unsafe.Pointer(&r))
	ctx, cmd, args := t.Cmd()
	fmt.Println(ctx, cmd, args)
	return "OK", nil
}
