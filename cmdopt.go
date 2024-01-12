package redis

import (
	"context"
	"unsafe"
)

type cmdOption interface {
	option(v ...any)
}

// optEx ...
type optEx[T cmdOption] struct{}

func (o *optEx[T]) Ex(seconds int64) T {
	t := *(*T)(unsafe.Pointer(&o))
	t.option("EX", seconds)
	return t
}

func (o *optEx[T]) Expire(seconds int64) T {
	t := *(*T)(unsafe.Pointer(&o))
	t.option("EX", seconds)
	return t
}

// PxOpt ...
type optPx[T cmdOption] struct{}

func (o *optPx[T]) Px(milliseconds int64) T {
	t := *(*T)(unsafe.Pointer(&o))
	t.option("PX", milliseconds)
	return t
}

func (o *optPx[T]) PExpire(milliseconds int64) T {
	t := *(*T)(unsafe.Pointer(&o))
	t.option("PX", milliseconds)
	return t
}

// optExAt ...
type optExAt[T cmdOption] struct{}

func (o *optExAt[T]) ExAt(timestamp int64) T {
	t := *(*T)(unsafe.Pointer(&o))
	t.option("EXAT", timestamp)
	return t
}

func (o *optExAt[T]) ExpireAt(timestamp int64) T {
	t := *(*T)(unsafe.Pointer(&o))
	t.option("EXAT", timestamp)
	return t
}

// PxAtOpt ...
type optPxAt[T cmdOption] struct{}

func (o *optPxAt[T]) PxAT(timestamp int64) T {
	t := *(*T)(unsafe.Pointer(&o))
	t.option("PXAT", timestamp)
	return t
}

func (o *optPxAt[T]) PExpireAT(timestamp int64) T {
	t := *(*T)(unsafe.Pointer(&o))
	t.option("PXAT", timestamp)
	return t
}

// optPersist ...
type optPersist[T cmdOption] struct{}

func (o *optPersist[T]) Persist() T {
	t := *(*T)(unsafe.Pointer(&o))
	t.option("PERSIST")
	return t
}

// optKeepTTL ...
type optKeepTTL[T cmdOption] struct{}

func (o *optKeepTTL[T]) KeepTTL() T {
	t := *(*T)(unsafe.Pointer(&o))
	t.option("KEEPTTL")
	return t
}

// optNX ...
type optNX[T cmdOption] struct{}

func (o *optNX[T]) Nx() T {
	t := *(*T)(unsafe.Pointer(&o))
	t.option("NX")
	return t
}

func (o *optNX[T]) NotExists() T {
	t := *(*T)(unsafe.Pointer(&o))
	t.option("NX")
	return t
}

// optXX ...
type optXX[T cmdOption] struct{}

func (o *optXX[T]) Xx() T {
	t := *(*T)(unsafe.Pointer(&o))
	t.option("XX")
	return t
}

func (o *optXX[T]) Exists() T {
	t := *(*T)(unsafe.Pointer(&o))
	t.option("XX")
	return t
}

// optGet ...
type optGet[T cmdOption] struct{}

func (o *optGet[T]) Get() T {
	t := *(*T)(unsafe.Pointer(&o))
	t.option("GET")
	return t
}

type command struct {
	d    Driver
	ctx  context.Context
	cmd  string
	args []any
}

func (impl *command) driver() Driver {
	return impl.d
}

func (impl *command) option(v ...any) {
	impl.args = append(impl.args, v...)
}

func (impl *command) cmdArgs() (ctx context.Context, cmd string, args []any) {
	return impl.ctx, impl.cmd, impl.args
}
