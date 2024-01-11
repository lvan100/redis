package redis

import "unsafe"

type CmdOption interface {
	AddOption(v ...any)
}

type CmdOptionImpl struct {
	options []any
}

func (impl *CmdOptionImpl) AddOption(v ...any) {
	impl.options = append(impl.options, v...)
}

type ExOpt[T CmdOption] struct{}

func (ex *ExOpt[T]) Ex(seconds int) T {
	t := *(*T)(unsafe.Pointer(&ex))
	t.AddOption("EX", seconds)
	return t
}
