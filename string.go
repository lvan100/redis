package redis

import (
	"bytes"
	"context"
	"fmt"
	"strconv"
)

func writeAny(buf *bytes.Buffer, v any) {
	buf.WriteString(fmt.Sprint(v))
}

func writeString(buf *bytes.Buffer, s string) {
	buf.WriteString(s)
}

type CmdOption func(*bytes.Buffer)

func Ex(seconds int) CmdOption {
	return func(buf *bytes.Buffer) {
		buf.WriteString("EX")
		buf.WriteByte(' ')
		buf.WriteString(strconv.Itoa(seconds))
		buf.WriteByte(' ')
	}
}

func Px(milliseconds int) CmdOption {
	return func(buf *bytes.Buffer) {

	}
}

func ExAt(timestamp int64) CmdOption {
	return func(buf *bytes.Buffer) {

	}
}

func PxAt(timestamp int64) CmdOption {
	return func(buf *bytes.Buffer) {

	}
}

func KeepTTL() CmdOption {
	return func(buf *bytes.Buffer) {
		buf.WriteString("KEEPTTL")
		buf.WriteByte(' ')
	}
}

func Nx() CmdOption {
	return func(buf *bytes.Buffer) {

	}
}

func Xx() CmdOption {
	return func(buf *bytes.Buffer) {

	}
}

func Get() CmdOption {
	return func(buf *bytes.Buffer) {

	}
}

type StringCmd interface {

	// Set https://redis.io/commands/set
	// Command: SET key value [EX seconds|PX milliseconds|EXAT timestamp|PXAT milliseconds-timestamp|KEEPTTL] [NX|XX] [GET]
	// Simple string reply: OK if SET was executed correctly.
	Set(ctx context.Context, key string, value any, opts ...CmdOption) (string, error)
}

type stringCmd struct {
	driver Driver
}

func (c *stringCmd) Set(ctx context.Context, key string, value any, opts ...CmdOption) (string, error) {
	buf := getBytes()
	defer putBytes(buf)
	writeString(buf, key)
	buf.WriteByte(' ')
	writeAny(buf, value)
	buf.WriteByte(' ')
	for _, opt := range opts {
		opt(buf)
	}
	r, err := c.driver.Exec(ctx, buf)
	if err != nil {
		return "", err
	}
	return r.(string), nil
}
