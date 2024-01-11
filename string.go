package redis

import (
	"context"
)

type StringCmdSet struct {
	StringReplier[*StringCmdSet]
	ExOpt[*StringCmdSet]
	CmdOptionImpl
	key string
	val any
	ctx context.Context
}

func (c *StringCmdSet) Cmd() (ctx context.Context, cmd string, args []any) {
	return c.ctx, "SET", append(append([]any{}, c.key, c.val), c.options...)
}

type StringCmd interface {

	// Set https://redis.io/commands/set
	// Command: SET key value [EX seconds|PX milliseconds|EXAT timestamp|PXAT milliseconds-timestamp|KEEPTTL] [NX|XX] [GET]
	// Simple string reply: OK if SET was executed correctly.
	Set(ctx context.Context, key string, value any) *StringCmdSet
}

type stringCmd struct {
	driver Driver
}

func (c *stringCmd) Set(ctx context.Context, key string, value any) *StringCmdSet {
	return &StringCmdSet{key: key, val: value, ctx: ctx}
}
