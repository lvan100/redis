package redis

import "context"

type HashOps interface {

	// HDel https://redis.io/commands/hdel
	// Command: HDEL key field [field ...]
	// Integer reply: the number of fields that were removed
	// from the hash, not including specified but non-existing fields.
	HDel(ctx context.Context, key string, fields ...string) *Int64Replier

	// HExists https://redis.io/commands/hexists
	// Command: HEXISTS key field
	// Integer reply: 1 if the hash contains field,
	// 0 if the hash does not contain field, or key does not exist.
	HExists(ctx context.Context, key, field string) *Int64Replier

	// HGet https://redis.io/commands/hget
	// Command: HGET key field
	// Bulk string reply: the value associated with field,
	// or nil when field is not present in the hash or key does not exist.
	HGet(ctx context.Context, key string, field string) *StringReplier

	// HGetAll https://redis.io/commands/hgetall
	// Command: HGETALL key
	// Array reply: list of fields and their values stored
	// in the hash, or an empty list when key does not exist.
	HGetAll(ctx context.Context, key string) *StringMapReplier

	// HIncrBy https://redis.io/commands/hincrby
	// Command: HINCRBY key field increment
	// Integer reply: the value at field after the increment operation.
	HIncrBy(ctx context.Context, key, field string, incr int64) *Int64Replier

	// HIncrByFloat https://redis.io/commands/hincrbyfloat
	// Command: HINCRBYFLOAT key field increment
	// Bulk string reply: the value of field after the increment.
	HIncrByFloat(ctx context.Context, key, field string, incr float64) *Float64Replier

	// HKeys https://redis.io/commands/hkeys
	// Command: HKEYS key
	// Array reply: list of fields in the hash, or an empty list when key does not exist.
	HKeys(ctx context.Context, key string) *StringSliceReplier

	// HLen https://redis.io/commands/hlen
	// Command: HLEN key
	// Integer reply: number of fields in the hash, or 0 when key does not exist.
	HLen(ctx context.Context, key string) *Int64Replier

	// HMGet https://redis.io/commands/hmget
	// Command: HMGET key field [field ...]
	// Array reply: list of values associated with the
	// given fields, in the same order as they are requested.
	HMGet(ctx context.Context, key string, fields ...string) *SliceReplier

	// HSet https://redis.io/commands/hset
	// Command: HSET key field value [field value ...]
	// Integer reply: The number of fields that were added.
	HSet(ctx context.Context, key string, args ...any) *Int64Replier

	// HSetNX https://redis.io/commands/hsetnx
	// Command: HSETNX key field value
	// Integer reply: 1 if field is a new field in the hash and value was set,
	// 0 if field already exists in the hash and no operation was performed.
	HSetNX(ctx context.Context, key, field string, value any) *Int64Replier

	// HStrLen https://redis.io/commands/hstrlen
	// Command: HSTRLEN key field
	// Integer reply: the string length of the value associated with field,
	// or zero when field is not present in the hash or key does not exist at all.
	HStrLen(ctx context.Context, key, field string) *Int64Replier

	// HVals https://redis.io/commands/hvals
	// Command: HVALS key
	// Array reply: list of values in the hash, or an empty list when key does not exist.
	HVals(ctx context.Context, key string) *StringSliceReplier
}

type hashOps struct {
	driver Driver
}

func (c *hashOps) HDel(ctx context.Context, key string, fields ...string) *Int64Replier {
	args := []any{key}
	for _, field := range fields {
		args = append(args, field)
	}
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "HDEL",
			args:   args,
		},
	}
}

func (c *hashOps) HExists(ctx context.Context, key, field string) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "HEXISTS",
			args:   []any{key, field},
		},
	}
}

func (c *hashOps) HGet(ctx context.Context, key string, field string) *StringReplier {
	return &StringReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "HGET",
			args:   []any{key, field},
		},
	}
}

func (c *hashOps) HGetAll(ctx context.Context, key string) *StringMapReplier {
	return &StringMapReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "HGETALL",
			args:   []any{key},
		},
	}
}

func (c *hashOps) HIncrBy(ctx context.Context, key, field string, incr int64) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "HINCRBY",
			args:   []any{key, field, incr},
		},
	}
}

func (c *hashOps) HIncrByFloat(ctx context.Context, key, field string, incr float64) *Float64Replier {
	return &Float64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "HINCRBYFLOAT",
			args:   []any{key, field, incr},
		},
	}
}

func (c *hashOps) HKeys(ctx context.Context, key string) *StringSliceReplier {
	return &StringSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "HKEYS",
			args:   []any{key},
		},
	}
}

func (c *hashOps) HLen(ctx context.Context, key string) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "HLEN",
			args:   []any{key},
		},
	}
}

func (c *hashOps) HMGet(ctx context.Context, key string, fields ...string) *SliceReplier {
	args := []any{key}
	for _, field := range fields {
		args = append(args, field)
	}
	return &SliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "HMGET",
			args:   args,
		},
	}
}

func (c *hashOps) HSet(ctx context.Context, key string, args ...any) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "HSET",
			args:   append([]any{key}, args),
		},
	}
}

func (c *hashOps) HSetNX(ctx context.Context, key, field string, value any) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "HSETNX",
			args:   []any{key, field, value},
		},
	}
}

func (c *hashOps) HStrLen(ctx context.Context, key, field string) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "HSTRLEN",
			args:   []any{key, field},
		},
	}
}

func (c *hashOps) HVals(ctx context.Context, key string) *StringSliceReplier {
	return &StringSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "HVALS",
			args:   []any{key},
		},
	}
}
