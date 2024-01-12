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
	CmdStringReplier
}

type CmdSet struct {
	optEx[*CmdSet]
	optPx[*CmdSet]
	optExAt[*CmdSet]
	optPxAt[*CmdSet]
	optKeepTTL[*CmdSet]
	optNX[*CmdSet]
	optXX[*CmdSet]
	optGet[*CmdSet]
	CmdStringReplier
}

type StringOps interface {

	// Append https://redis.io/commands/append
	// Command: APPEND key value
	// Integer reply: the length of the string after the append operation.
	Append(ctx context.Context, key, value string) *CmdIntReplier

	// Decr https://redis.io/commands/decr
	// Command: DECR key
	// Integer reply: the value of key after the decrement
	Decr(ctx context.Context, key string) (int64, error)

	// DecrBy https://redis.io/commands/decrby
	// Command: DECRBY key decrement
	// Integer reply: the value of key after the decrement.
	DecrBy(ctx context.Context, key string, decrement int64) (int64, error)

	// Get https://redis.io/commands/get
	// Command: GET key
	// Bulk string reply: the value of key, or nil when key does not exist.
	Get(ctx context.Context, key string) (string, error)

	// GetDel https://redis.io/commands/getdel
	// Command: GETDEL key
	// Bulk string reply: the value of key, nil when key does not exist,
	// or an error if the key's value type isn't a string.
	GetDel(ctx context.Context, key string) (string, error)

	// GetEx https://redis.io/commands/getex
	// Command: GETEX key [EX seconds|PX milliseconds|EXAT timestamp|PXAT milliseconds-timestamp|PERSIST]
	// Bulk string reply: the value of key, or nil when key does not exist.
	GetEx(ctx context.Context, key string) *CmdGet

	// GetRange https://redis.io/commands/getrange
	// Command: GETRANGE key start end
	// Bulk string reply
	GetRange(ctx context.Context, key string, start, end int64) (string, error)

	// GetSet https://redis.io/commands/getset
	// Command: GETSET key value
	// Bulk string reply: the old value stored at key, or nil when key did not exist.
	GetSet(ctx context.Context, key string, value any) (string, error)

	// Incr https://redis.io/commands/incr
	// Command: INCR key
	// Integer reply: the value of key after the increment
	Incr(ctx context.Context, key string) (int64, error)

	// IncrBy https://redis.io/commands/incrby
	// Command: INCRBY key increment
	// Integer reply: the value of key after the increment.
	IncrBy(ctx context.Context, key string, value int64) (int64, error)

	// IncrByFloat https://redis.io/commands/incrbyfloat
	// Command: INCRBYFLOAT key increment
	// Bulk string reply: the value of key after the increment.
	IncrByFloat(ctx context.Context, key string, value float64) (float64, error)

	// MGet https://redis.io/commands/mget
	// Command: MGET key [key ...]
	// Array reply: list of values at the specified keys.
	MGet(ctx context.Context, keys ...string) ([]any, error)

	// MSet https://redis.io/commands/mset
	// Command: MSET key value [key value ...]
	// Simple string reply: always OK since MSET can't fail.
	MSet(ctx context.Context, args ...any) (string, error)

	// MSetNX https://redis.io/commands/msetnx
	// Command: MSETNX key value [key value ...]
	// MSETNX is atomic, so all given keys are set at once
	// Integer reply: 1 if the all the keys were set, 0 if no
	// key was set (at least one key already existed).
	MSetNX(ctx context.Context, args ...any) (int64, error)

	// PSetEX https://redis.io/commands/psetex
	// Command: PSETEX key milliseconds value
	// Simple string reply
	PSetEX(ctx context.Context, key string, value any, expire int64) (string, error)

	// Set https://redis.io/commands/set
	// Command: SET key value [EX seconds|PX milliseconds|EXAT timestamp|PXAT milliseconds-timestamp|KEEPTTL] [NX|XX] [GET]
	// Simple string reply: OK if SET was executed correctly.
	Set(ctx context.Context, key string, value any) *CmdSet

	// SetEX https://redis.io/commands/setex
	// Command: SETEX key seconds value
	// Simple string reply
	SetEX(ctx context.Context, key string, value any, expire int64) (string, error)

	// SetNX https://redis.io/commands/setnx
	// Command: SETNX key value
	// Integer reply: 1 if the key was set, 0 if the key was not set.
	SetNX(ctx context.Context, key string, value any) (int64, error)

	// SetRange https://redis.io/commands/setrange
	// Command: SETRANGE key offset value
	// Integer reply: the length of the string after it was modified by the command.
	SetRange(ctx context.Context, key string, offset int64, value string) (int64, error)

	// StrLen https://redis.io/commands/strlen
	// Command: STRLEN key
	// Integer reply: the length of the string at key, or 0 when key does not exist.
	StrLen(ctx context.Context, key string) (int64, error)
}

var _ StringOps = (*stringOps)(nil)

type stringOps struct {
	driver Driver
}

func (c *stringOps) Append(ctx context.Context, key, value string) *CmdIntReplier {
	return &CmdIntReplier{command: command{
		d:    c.driver,
		ctx:  ctx,
		cmd:  "APPEND",
		args: []any{key, value},
	}}
}

func (c *stringOps) Decr(ctx context.Context, key string) (int64, error) {
	return toInt64(c.driver.Exec(ctx, "DECR", []any{key}))
}

func (c *stringOps) DecrBy(ctx context.Context, key string, decrement int64) (int64, error) {
	return toInt64(c.driver.Exec(ctx, "DECRBY", []any{key, decrement}))
}

func (c *stringOps) Get(ctx context.Context, key string) (string, error) {
	return toString(c.driver.Exec(ctx, "GET", []any{key}))
}

func (c *stringOps) GetDel(ctx context.Context, key string) (string, error) {
	return toString(c.driver.Exec(ctx, "GETDEL", []any{key}))
}

func (c *stringOps) GetEx(ctx context.Context, key string) *CmdGet {
	return &CmdGet{
		CmdStringReplier: CmdStringReplier{
			command: command{
				d:    c.driver,
				ctx:  ctx,
				cmd:  "GETEX",
				args: []any{key},
			},
		},
	}
}

func (c *stringOps) GetRange(ctx context.Context, key string, start, end int64) (string, error) {
	return toString(c.driver.Exec(ctx, "GETRANGE", []any{key, start, end}))
}

func (c *stringOps) GetSet(ctx context.Context, key string, value any) (string, error) {
	return toString(c.driver.Exec(ctx, "GETSET", []any{key, value}))
}

func (c *stringOps) Incr(ctx context.Context, key string) (int64, error) {
	return toInt64(c.driver.Exec(ctx, "INCR", []any{key}))
}

func (c *stringOps) IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	return toInt64(c.driver.Exec(ctx, "INCRBY", []any{key, value}))
}

func (c *stringOps) IncrByFloat(ctx context.Context, key string, value float64) (float64, error) {
	return toFloat64(c.driver.Exec(ctx, "INCRBYFLOAT", []any{key, value}))
}

func (c *stringOps) MGet(ctx context.Context, keys ...string) ([]any, error) {
	args := make([]any, 0, len(keys))
	for _, key := range keys {
		args = append(args, key)
	}
	return toSlice(c.driver.Exec(ctx, "MGET", args))
}

func (c *stringOps) MSet(ctx context.Context, args ...any) (string, error) {
	return toString(c.driver.Exec(ctx, "MSET", args))
}

func (c *stringOps) MSetNX(ctx context.Context, args ...any) (int64, error) {
	return toInt64(c.driver.Exec(ctx, "MSETNX", args))
}

func (c *stringOps) PSetEX(ctx context.Context, key string, value any, expire int64) (string, error) {
	return toString(c.driver.Exec(ctx, "PSETEX", []any{key, expire, value}))
}

func (c *stringOps) Set(ctx context.Context, key string, value any) *CmdSet {
	return &CmdSet{
		CmdStringReplier: CmdStringReplier{
			command: command{
				d:    c.driver,
				ctx:  ctx,
				cmd:  "SET",
				args: []any{key, value},
			},
		},
	}
}

func (c *stringOps) SetEX(ctx context.Context, key string, value any, expire int64) (string, error) {
	return toString(c.driver.Exec(ctx, "SETEX", []any{key, expire, value}))
}

func (c *stringOps) SetNX(ctx context.Context, key string, value any) (int64, error) {
	return toInt64(c.driver.Exec(ctx, "SETNX", []any{key, value}))
}

func (c *stringOps) SetRange(ctx context.Context, key string, offset int64, value string) (int64, error) {
	return toInt64(c.driver.Exec(ctx, "SETRANGE", []any{key, offset, value}))
}

func (c *stringOps) StrLen(ctx context.Context, key string) (int64, error) {
	return toInt64(c.driver.Exec(ctx, "STRLEN", []any{key}))
}
