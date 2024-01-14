package redis

import "context"

type CmdExpire struct {
	optNX[*CmdExpire]
	optXX[*CmdExpire]
	optGT[*CmdExpire]
	optLT[*CmdExpire]
	Int64Replier
}

type CmdExpireAt struct {
	optNX[*CmdExpireAt]
	optXX[*CmdExpireAt]
	optGT[*CmdExpireAt]
	optLT[*CmdExpireAt]
	Int64Replier
}

type CmdPExpire struct {
	optNX[*CmdPExpire]
	optXX[*CmdPExpire]
	optGT[*CmdPExpire]
	optLT[*CmdPExpire]
	Int64Replier
}

type CmdPExpireAt struct {
	optNX[*CmdPExpireAt]
	optXX[*CmdPExpireAt]
	optGT[*CmdPExpireAt]
	optLT[*CmdPExpireAt]
	Int64Replier
}

type KeyOps interface {

	// Del https://redis.io/commands/del
	// Command: DEL key [key ...]
	// Integer reply: The number of keys that were removed.
	Del(ctx context.Context, keys ...string) *Int64Replier

	// Dump https://redis.io/commands/dump
	// Command: DUMP key
	// Bulk string reply: the serialized value.
	// If key does not exist a nil bulk reply is returned.
	Dump(ctx context.Context, key string) *StringReplier

	// Exists https://redis.io/commands/exists
	// Command: EXISTS key [key ...]
	// Integer reply: The number of keys existing among the
	// ones specified as arguments. Keys mentioned multiple
	// times and existing are counted multiple times.
	Exists(ctx context.Context, keys ...string) *Int64Replier

	// Expire https://redis.io/commands/expire
	// Command: EXPIRE key seconds [NX|XX|GT|LT]
	// Integer reply: 1 if the timeout was set, 0 if the timeout was not set.
	Expire(ctx context.Context, key string, expire int64) *CmdExpire

	// ExpireAt https://redis.io/commands/expireat
	// Command: EXPIREAT key timestamp [NX|XX|GT|LT]
	// Integer reply: 1 if the timeout was set, 0 if the timeout was not set.
	ExpireAt(ctx context.Context, key string, expireAt int64) *CmdExpireAt

	// Keys https://redis.io/commands/keys
	// Command: KEYS pattern
	// Array reply: list of keys matching pattern.
	Keys(ctx context.Context, pattern string) *StringSliceReplier

	// Persist https://redis.io/commands/persist
	// Command: PERSIST key
	// Integer reply: 1 if the timeout was removed,
	// 0 if key does not exist or does not have an associated timeout.
	Persist(ctx context.Context, key string) *Int64Replier

	// PExpire https://redis.io/commands/pexpire
	// Command: PEXPIRE key milliseconds [NX|XX|GT|LT]
	// Integer reply: 1 if the timeout was set, 0 if the timeout was not set.
	PExpire(ctx context.Context, key string, expire int64) *CmdPExpire

	// PExpireAt https://redis.io/commands/pexpireat
	// Command: PEXPIREAT key milliseconds-timestamp [NX|XX|GT|LT]
	// Integer reply: 1 if the timeout was set, 0 if the timeout was not set.
	PExpireAt(ctx context.Context, key string, expireAt int64) *CmdPExpireAt

	// PTTL https://redis.io/commands/pttl
	// Command: PTTL key
	// Integer reply: TTL in milliseconds, -1 if the key exists
	// but has no associated expire, -2 if the key does not exist.
	PTTL(ctx context.Context, key string) *Int64Replier

	// RandomKey https://redis.io/commands/randomkey
	// Command: RANDOMKEY
	// Bulk string reply: the random key, or nil when the database is empty.
	RandomKey(ctx context.Context) *StringReplier

	// Rename https://redis.io/commands/rename
	// Command: RENAME key newkey
	// Simple string reply.
	Rename(ctx context.Context, key, newKey string) *StringReplier

	// RenameNX https://redis.io/commands/renamenx
	// Command: RENAMENX key newkey
	// Integer reply: 1 if key was renamed to newKey, 0 if newKey already exists.
	RenameNX(ctx context.Context, key, newKey string) *Int64Replier

	// Touch https://redis.io/commands/touch
	// Command: TOUCH key [key ...]
	// Integer reply: The number of keys that were touched.
	Touch(ctx context.Context, keys ...string) *Int64Replier

	// TTL https://redis.io/commands/ttl
	// Command: TTL key
	// Integer reply: TTL in seconds, -1 if the key exists
	// but has no associated expire, -2 if the key does not exist.
	TTL(ctx context.Context, key string) *Int64Replier

	// Type https://redis.io/commands/type
	// Command: TYPE key
	// Simple string reply: type of key, or none when key does not exist.
	Type(ctx context.Context, key string) *StringReplier
}

///////////////////////

var _ KeyOps = (*keyOps)(nil)

type keyOps struct {
	driver Driver
}

func (c *keyOps) Del(ctx context.Context, keys ...string) *Int64Replier {
	var args []any
	for _, key := range keys {
		args = append(args, key)
	}
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "DEL",
			args:   args,
		},
	}
}

func (c *keyOps) Dump(ctx context.Context, key string) *StringReplier {
	return &StringReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "DUMP",
			args:   []any{key},
		},
	}
}

func (c *keyOps) Exists(ctx context.Context, keys ...string) *Int64Replier {
	var args []any
	for _, key := range keys {
		args = append(args, key)
	}
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "EXISTS",
			args:   args,
		},
	}
}

func (c *keyOps) Expire(ctx context.Context, key string, expire int64) *CmdExpire {
	return &CmdExpire{
		Int64Replier: Int64Replier{
			command: command{
				driver: c.driver,
				ctx:    ctx,
				cmd:    "EXPIRE",
				args:   []any{key, expire},
			},
		},
	}
}

func (c *keyOps) ExpireAt(ctx context.Context, key string, expireAt int64) *CmdExpireAt {
	return &CmdExpireAt{
		Int64Replier: Int64Replier{
			command: command{
				driver: c.driver,
				ctx:    ctx,
				cmd:    "EXPIREAT",
				args:   []any{key, expireAt},
			},
		},
	}
}

func (c *keyOps) Keys(ctx context.Context, pattern string) *StringSliceReplier {
	return &StringSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "KEYS",
			args:   []any{pattern},
		},
	}
}

func (c *keyOps) Persist(ctx context.Context, key string) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "PERSIST",
			args:   []any{key},
		},
	}
}

func (c *keyOps) PExpire(ctx context.Context, key string, expire int64) *CmdPExpire {
	return &CmdPExpire{
		Int64Replier: Int64Replier{
			command: command{
				driver: c.driver,
				ctx:    ctx,
				cmd:    "PEXPIRE",
				args:   []any{key, expire},
			},
		},
	}
}

func (c *keyOps) PExpireAt(ctx context.Context, key string, expireAt int64) *CmdPExpireAt {
	return &CmdPExpireAt{
		Int64Replier: Int64Replier{
			command: command{
				driver: c.driver,
				ctx:    ctx,
				cmd:    "PEXPIREAT",
				args:   []any{key, expireAt},
			},
		},
	}
}

func (c *keyOps) PTTL(ctx context.Context, key string) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "PTTL",
			args:   []any{key},
		},
	}
}

func (c *keyOps) RandomKey(ctx context.Context) *StringReplier {
	return &StringReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "RANDOMKEY",
			args:   nil,
		},
	}
}

func (c *keyOps) Rename(ctx context.Context, key, newKey string) *StringReplier {
	return &StringReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "RENAME",
			args:   []any{key, newKey},
		},
	}
}

func (c *keyOps) RenameNX(ctx context.Context, key, newKey string) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "RENAMENX",
			args:   []any{key, newKey},
		},
	}
}

func (c *keyOps) Touch(ctx context.Context, keys ...string) *Int64Replier {
	var args []any
	for _, key := range keys {
		args = append(args, key)
	}
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "TOUCH",
			args:   args,
		},
	}
}

func (c *keyOps) TTL(ctx context.Context, key string) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "TTL",
			args:   []any{key},
		},
	}
}

func (c *keyOps) Type(ctx context.Context, key string) *StringReplier {
	return &StringReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "TYPE",
			args:   []any{key},
		},
	}
}
