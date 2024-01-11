package redis

import (
	"context"
	"unsafe"
)

type CmdOption interface {
	Option(v ...any)
}

type ExOpt[T CmdOption] struct{}

func (ex *ExOpt[T]) Ex(seconds int) T {
	t := *(*T)(unsafe.Pointer(&ex))
	t.Option("EX", seconds)
	return t
}

type Command struct {
	driver Driver
	ctx    context.Context
	cmd    string
	args   []any
}

func (impl *Command) Driver() Driver {
	return impl.driver
}

func (impl *Command) Option(v ...any) {
	impl.args = append(impl.args, v...)
}

func (impl *Command) CmdArgs() (ctx context.Context, cmd string, args []any) {
	return impl.ctx, "SET", impl.args
}
