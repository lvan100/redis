package redis

import (
	"context"
)

type CmdGet struct {
	optEx[*CmdGet]
	optPx[*CmdGet]
	optExAt[*CmdGet]
	optPxAt[*CmdGet]
	optPersist[*CmdGet]
	StringReplier
}

type CmdSet struct {
	optEx[*CmdSet]
	optPx[*CmdSet]
	optExAt[*CmdSet]
	optPxAt[*CmdSet]
	optNX[*CmdSet]
	optXX[*CmdSet]
	optKeepTTL[*CmdSet]
	optGet[*CmdSet]
	StringReplier
}

type StringOps interface {

	// Append https://redis.io/commands/append
	// Command: APPEND key value
	// Integer reply: the length of the string after the append operation.
	Append(ctx context.Context, key, value string) *Int64Replier

	// Decr https://redis.io/commands/decr
	// Command: DECR key
	// Integer reply: the value of key after the decrement
	Decr(ctx context.Context, key string) *Int64Replier

	// DecrBy https://redis.io/commands/decrby
	// Command: DECRBY key decrement
	// Integer reply: the value of key after the decrement.
	DecrBy(ctx context.Context, key string, decrement int64) *Int64Replier

	// Get https://redis.io/commands/get
	// Command: GET key
	// Bulk string reply: the value of key, or nil when key does not exist.
	Get(ctx context.Context, key string) *StringReplier

	// GetDel https://redis.io/commands/getdel
	// Command: GETDEL key
	// Bulk string reply: the value of key, nil when key does not exist,
	// or an error if the key's value type isn't a string.
	GetDel(ctx context.Context, key string) *StringReplier

	// GetEx https://redis.io/commands/getex
	// Command: GETEX key [EX seconds|PX milliseconds|EXAT timestamp|PXAT milliseconds-timestamp|PERSIST]
	// Bulk string reply: the value of key, or nil when key does not exist.
	GetEx(ctx context.Context, key string) *CmdGet

	// GetRange https://redis.io/commands/getrange
	// Command: GETRANGE key start end
	// Bulk string reply
	GetRange(ctx context.Context, key string, start, end int64) *StringReplier

	// GetSet https://redis.io/commands/getset
	// Command: GETSET key value
	// Bulk string reply: the old value stored at key, or nil when key did not exist.
	GetSet(ctx context.Context, key string, value any) *StringReplier

	// Incr https://redis.io/commands/incr
	// Command: INCR key
	// Integer reply: the value of key after the increment
	Incr(ctx context.Context, key string) *Int64Replier

	// IncrBy https://redis.io/commands/incrby
	// Command: INCRBY key increment
	// Integer reply: the value of key after the increment.
	IncrBy(ctx context.Context, key string, value int64) *Int64Replier

	// IncrByFloat https://redis.io/commands/incrbyfloat
	// Command: INCRBYFLOAT key increment
	// Bulk string reply: the value of key after the increment.
	IncrByFloat(ctx context.Context, key string, value float64) *Float64Replier

	// MGet https://redis.io/commands/mget
	// Command: MGET key [key ...]
	// Array reply: list of values at the specified keys.
	MGet(ctx context.Context, keys ...string) *SliceReplier

	// MSet https://redis.io/commands/mset
	// Command: MSET key value [key value ...]
	// Simple string reply: always OK since MSET can't fail.
	MSet(ctx context.Context, args ...any) *StringReplier

	// MSetNX https://redis.io/commands/msetnx
	// Command: MSETNX key value [key value ...]
	// MSETNX is atomic, so all given keys are set at once
	// Integer reply: 1 if the all the keys were set, 0 if no
	// key was set (at least one key already existed).
	MSetNX(ctx context.Context, args ...any) *Int64Replier

	// PSetEX https://redis.io/commands/psetex
	// Command: PSETEX key milliseconds value
	// Simple string reply
	PSetEX(ctx context.Context, key string, value any, expire int64) *StringReplier

	// Set https://redis.io/commands/set
	// Command: SET key value [EX seconds|PX milliseconds|EXAT timestamp|PXAT milliseconds-timestamp|KEEPTTL] [NX|XX] [GET]
	// Simple string reply: OK if SET was executed correctly.
	Set(ctx context.Context, key string, value any) *CmdSet

	// SetEX https://redis.io/commands/setex
	// Command: SETEX key seconds value
	// Simple string reply
	SetEX(ctx context.Context, key string, value any, expire int64) *StringReplier

	// SetNX https://redis.io/commands/setnx
	// Command: SETNX key value
	// Integer reply: 1 if the key was set, 0 if the key was not set.
	SetNX(ctx context.Context, key string, value any) *Int64Replier

	// SetRange https://redis.io/commands/setrange
	// Command: SETRANGE key offset value
	// Integer reply: the length of the string after it was modified by the command.
	SetRange(ctx context.Context, key string, offset int64, value string) *Int64Replier

	// StrLen https://redis.io/commands/strlen
	// Command: STRLEN key
	// Integer reply: the length of the string at key, or 0 when key does not exist.
	StrLen(ctx context.Context, key string) *Int64Replier
}

var _ StringOps = (*stringOps)(nil)

type stringOps struct {
	driver Driver
}

func (c *stringOps) Append(ctx context.Context, key, value string) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "APPEND",
			args:   []any{key, value},
		},
	}
}

func (c *stringOps) Decr(ctx context.Context, key string) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "DECR",
			args:   []any{key},
		},
	}
}

func (c *stringOps) DecrBy(ctx context.Context, key string, decrement int64) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "DECRBY",
			args:   []any{key, decrement},
		},
	}
}

func (c *stringOps) Get(ctx context.Context, key string) *StringReplier {
	return &StringReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "GET",
			args:   []any{key},
		},
	}
}

func (c *stringOps) GetDel(ctx context.Context, key string) *StringReplier {
	return &StringReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "GETDEL",
			args:   []any{key},
		},
	}
}

func (c *stringOps) GetEx(ctx context.Context, key string) *CmdGet {
	return &CmdGet{
		StringReplier: StringReplier{
			command: command{
				driver: c.driver,
				ctx:    ctx,
				cmd:    "GETEX",
				args:   []any{key},
			},
		},
	}
}

func (c *stringOps) GetRange(ctx context.Context, key string, start, end int64) *StringReplier {
	return &StringReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "GETRANGE",
			args:   []any{key, start, end},
		},
	}
}

func (c *stringOps) GetSet(ctx context.Context, key string, value any) *StringReplier {
	return &StringReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "GETSET",
			args:   []any{key, value},
		},
	}
}

func (c *stringOps) Incr(ctx context.Context, key string) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "INCR",
			args:   []any{key},
		},
	}
}

func (c *stringOps) IncrBy(ctx context.Context, key string, value int64) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "INCRBY",
			args:   []any{key, value},
		},
	}
}

func (c *stringOps) IncrByFloat(ctx context.Context, key string, value float64) *Float64Replier {
	return &Float64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "INCRBYFLOAT",
			args:   []any{key, value},
		},
	}
}

func (c *stringOps) MGet(ctx context.Context, keys ...string) *SliceReplier {
	args := make([]any, 0, len(keys))
	for _, key := range keys {
		args = append(args, key)
	}
	return &SliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "MGET",
			args:   args,
		},
	}
}

func (c *stringOps) MSet(ctx context.Context, args ...any) *StringReplier {
	return &StringReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "MSET",
			args:   args,
		},
	}
}

func (c *stringOps) MSetNX(ctx context.Context, args ...any) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "MSETNX",
			args:   args,
		},
	}
}

func (c *stringOps) PSetEX(ctx context.Context, key string, value any, expire int64) *StringReplier {
	return &StringReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "PSETEX",
			args:   []any{key, expire, value},
		},
	}
}

func (c *stringOps) Set(ctx context.Context, key string, value any) *CmdSet {
	return &CmdSet{
		StringReplier: StringReplier{
			command: command{
				driver: c.driver,
				ctx:    ctx,
				cmd:    "SET",
				args:   []any{key, value},
			},
		},
	}
}

func (c *stringOps) SetEX(ctx context.Context, key string, value any, expire int64) *StringReplier {
	return &StringReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "SETEX",
			args:   []any{key, expire, value},
		},
	}
}

func (c *stringOps) SetNX(ctx context.Context, key string, value any) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "SETNX",
			args:   []any{key, value},
		},
	}
}

func (c *stringOps) SetRange(ctx context.Context, key string, offset int64, value string) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "SETRANGE",
			args:   []any{key, offset, value},
		},
	}
}

func (c *stringOps) StrLen(ctx context.Context, key string) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "STRLEN",
			args:   []any{key},
		},
	}
}
