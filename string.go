package redis

import (
	"context"
)

type StringCmdSet struct {
	StringReplier[*StringCmdSet]
	ExOpt[*StringCmdSet]
	Command
}

type StringCmd interface {

	// Set https://redis.io/commands/set
	// Command: SET key value [EX seconds|PX milliseconds|EXAT timestamp|PXAT milliseconds-timestamp|KEEPTTL] [NX|XX] [GET]
	// Simple string reply: OK if SET was executed correctly.
	Set(ctx context.Context, key string, value any) *StringCmdSet
	//
	//// Append https://redis.io/commands/append
	//// Command: APPEND key value
	//// Integer reply: the length of the string after the append operation.
	//Append(ctx context.Context, key, value string) (int64, error)
	//
	//// Decr https://redis.io/commands/decr
	//// Command: DECR key
	//// Integer reply: the value of key after the decrement
	//Decr(ctx context.Context, key string) (int64, error)
	//
	//// DecrBy https://redis.io/commands/decrby
	//// Command: DECRBY key decrement
	//// Integer reply: the value of key after the decrement.
	//DecrBy(ctx context.Context, key string, decrement int64) (int64, error)
	//
	//// Get https://redis.io/commands/get
	//// Command: GET key
	//// Bulk string reply: the value of key, or nil when key does not exist.
	//Get(ctx context.Context, key string) (string, error)
	//
	//// GetDel https://redis.io/commands/getdel
	//// Command: GETDEL key
	//// Bulk string reply: the value of key, nil when key does not exist,
	//// or an error if the key's value type isn't a string.
	//GetDel(ctx context.Context, key string) (string, error)
	//
	//// GetEx https://redis.io/commands/getex
	//// Command: GETEX key [EX seconds|PX milliseconds|EXAT timestamp|PXAT milliseconds-timestamp|PERSIST]
	//// Bulk string reply: the value of key, or nil when key does not exist.
	//GetEx(ctx context.Context, key string, args ...any) (string, error)
	//
	//// GetRange https://redis.io/commands/getrange
	//// Command: GETRANGE key start end
	//// Bulk string reply
	//GetRange(ctx context.Context, key string, start, end int64) (string, error)
	//
	//// GetSet https://redis.io/commands/getset
	//// Command: GETSET key value
	//// Bulk string reply: the old value stored at key, or nil when key did not exist.
	//GetSet(ctx context.Context, key string, value any) (string, error)
	//
	//// Incr https://redis.io/commands/incr
	//// Command: INCR key
	//// Integer reply: the value of key after the increment
	//Incr(ctx context.Context, key string) (int64, error)
	//
	//// IncrBy https://redis.io/commands/incrby
	//// Command: INCRBY key increment
	//// Integer reply: the value of key after the increment.
	//IncrBy(ctx context.Context, key string, value int64) (int64, error)
	//
	//// IncrByFloat https://redis.io/commands/incrbyfloat
	//// Command: INCRBYFLOAT key increment
	//// Bulk string reply: the value of key after the increment.
	//IncrByFloat(ctx context.Context, key string, value float64) (float64, error)
	//
	//// MGet https://redis.io/commands/mget
	//// Command: MGET key [key ...]
	//// Array reply: list of values at the specified keys.
	//MGet(ctx context.Context, keys ...string) ([]any, error)
	//
	//// MSet https://redis.io/commands/mset
	//// Command: MSET key value [key value ...]
	//// Simple string reply: always OK since MSET can't fail.
	//MSet(ctx context.Context, args ...any) (string, error)
	//
	//// MSetNX https://redis.io/commands/msetnx
	//// Command: MSETNX key value [key value ...]
	//// MSETNX is atomic, so all given keys are set at once
	//// Integer reply: 1 if the all the keys were set, 0 if no
	//// key was set (at least one key already existed).
	//MSetNX(ctx context.Context, args ...any) (int64, error)
	//
	//// PSetEX https://redis.io/commands/psetex
	//// Command: PSETEX key milliseconds value
	//// Simple string reply
	//PSetEX(ctx context.Context, key string, value any, expire int64) (string, error)
	//
	//// Set https://redis.io/commands/set
	//// Command: SET key value [EX seconds|PX milliseconds|EXAT timestamp|PXAT milliseconds-timestamp|KEEPTTL] [NX|XX] [GET]
	//// Simple string reply: OK if SET was executed correctly.
	//Set(ctx context.Context, key string, value any, args ...any) (string, error)
	//
	//// SetEX https://redis.io/commands/setex
	//// Command: SETEX key seconds value
	//// Simple string reply
	//SetEX(ctx context.Context, key string, value any, expire int64) (string, error)
	//
	//// SetNX https://redis.io/commands/setnx
	//// Command: SETNX key value
	//// Integer reply: 1 if the key was set, 0 if the key was not set.
	//SetNX(ctx context.Context, key string, value any) (int64, error)
	//
	//// SetRange https://redis.io/commands/setrange
	//// Command: SETRANGE key offset value
	//// Integer reply: the length of the string after it was modified by the command.
	//SetRange(ctx context.Context, key string, offset int64, value string) (int64, error)
	//
	//// StrLen https://redis.io/commands/strlen
	//// Command: STRLEN key
	//// Integer reply: the length of the string at key, or 0 when key does not exist.
	//StrLen(ctx context.Context, key string) (int64, error)
}

type stringCmd struct {
	driver Driver
}

func (c *stringCmd) Set(ctx context.Context, key string, value any) *StringCmdSet {
	return &StringCmdSet{Command: Command{
		driver: c.driver,
		ctx:    ctx,
		cmd:    "SET",
		args:   append([]any{}, key, value),
	}}
}
