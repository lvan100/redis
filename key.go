package redis

import "context"

type KeyOps interface {

	// Del https://redis.io/commands/del
	// Command: DEL key [key ...]
	// Integer reply: The number of keys that were removed.
	Del(ctx context.Context, keys ...string) (int64, error)

	// Dump https://redis.io/commands/dump
	// Command: DUMP key
	// Bulk string reply: the serialized value.
	// If key does not exist a nil bulk reply is returned.
	Dump(ctx context.Context, key string) (string, error)

	// Exists https://redis.io/commands/exists
	// Command: EXISTS key [key ...]
	// Integer reply: The number of keys existing among the
	// ones specified as arguments. Keys mentioned multiple
	// times and existing are counted multiple times.
	Exists(ctx context.Context, keys ...string) (int64, error)

	// Expire https://redis.io/commands/expire
	// Command: EXPIRE key seconds [NX|XX|GT|LT]
	// Integer reply: 1 if the timeout was set, 0 if the timeout was not set.
	Expire(ctx context.Context, key string, expire int64, args ...any) (int64, error)

	// ExpireAt https://redis.io/commands/expireat
	// Command: EXPIREAT key timestamp [NX|XX|GT|LT]
	// Integer reply: 1 if the timeout was set, 0 if the timeout was not set.
	ExpireAt(ctx context.Context, key string, expireAt int64, args ...any) (int64, error)

	// Keys https://redis.io/commands/keys
	// Command: KEYS pattern
	// Array reply: list of keys matching pattern.
	Keys(ctx context.Context, pattern string) ([]string, error)

	// Persist https://redis.io/commands/persist
	// Command: PERSIST key
	// Integer reply: 1 if the timeout was removed,
	// 0 if key does not exist or does not have an associated timeout.
	Persist(ctx context.Context, key string) (int64, error)

	// PExpire https://redis.io/commands/pexpire
	// Command: PEXPIRE key milliseconds [NX|XX|GT|LT]
	// Integer reply: 1 if the timeout was set, 0 if the timeout was not set.
	PExpire(ctx context.Context, key string, expire int64, args ...any) (int64, error)

	// PExpireAt https://redis.io/commands/pexpireat
	// Command: PEXPIREAT key milliseconds-timestamp [NX|XX|GT|LT]
	// Integer reply: 1 if the timeout was set, 0 if the timeout was not set.
	PExpireAt(ctx context.Context, key string, expireAt int64, args ...any) (int64, error)

	// PTTL https://redis.io/commands/pttl
	// Command: PTTL key
	// Integer reply: TTL in milliseconds, -1 if the key exists
	// but has no associated expire, -2 if the key does not exist.
	PTTL(ctx context.Context, key string) (int64, error)

	// RandomKey https://redis.io/commands/randomkey
	// Command: RANDOMKEY
	// Bulk string reply: the random key, or nil when the database is empty.
	RandomKey(ctx context.Context) (string, error)

	// Rename https://redis.io/commands/rename
	// Command: RENAME key newkey
	// Simple string reply.
	Rename(ctx context.Context, key, newKey string) (string, error)

	// RenameNX https://redis.io/commands/renamenx
	// Command: RENAMENX key newkey
	// Integer reply: 1 if key was renamed to newKey, 0 if newKey already exists.
	RenameNX(ctx context.Context, key, newKey string) (int64, error)

	// Touch https://redis.io/commands/touch
	// Command: TOUCH key [key ...]
	// Integer reply: The number of keys that were touched.
	Touch(ctx context.Context, keys ...string) (int64, error)

	// TTL https://redis.io/commands/ttl
	// Command: TTL key
	// Integer reply: TTL in seconds, -1 if the key exists
	// but has no associated expire, -2 if the key does not exist.
	TTL(ctx context.Context, key string) (int64, error)

	// Type https://redis.io/commands/type
	// Command: TYPE key
	// Simple string reply: type of key, or none when key does not exist.
	Type(ctx context.Context, key string) (string, error)
}

///////////////////////

var _ KeyOps = (*keyOps)(nil)

type keyOps struct {
	driver Driver
}

// Del https://redis.io/commands/del
// Command: DEL key [key ...]
// Integer reply: The number of keys that were removed.
func (c *keyOps) Del(ctx context.Context, keys ...string) (int64, error) {
	var args []any
	for _, key := range keys {
		args = append(args, key)
	}
	return toInt64(c.driver.Exec(ctx, "DEL", args))
}

// Dump https://redis.io/commands/dump
// Command: DUMP key
// Bulk string reply: the serialized value.
// If key does not exist a nil bulk reply is returned.
func (c *keyOps) Dump(ctx context.Context, key string) (string, error) {
	args := []any{key}
	return toString(c.driver.Exec(ctx, "DUMP", args))
}

// Exists https://redis.io/commands/exists
// Command: EXISTS key [key ...]
// Integer reply: The number of keys existing among the
// ones specified as arguments. Keys mentioned multiple
// times and existing are counted multiple times.
func (c *keyOps) Exists(ctx context.Context, keys ...string) (int64, error) {
	args := []any{}
	for _, key := range keys {
		args = append(args, key)
	}
	return toInt64(c.driver.Exec(ctx, "EXISTS", args))
}

// Expire https://redis.io/commands/expire
// Command: EXPIRE key seconds [NX|XX|GT|LT]
// Integer reply: 1 if the timeout was set, 0 if the timeout was not set.
func (c *keyOps) Expire(ctx context.Context, key string, expire int64, args ...any) (int64, error) {
	args = append([]any{key, expire}, args)
	return toInt64(c.driver.Exec(ctx, "EXPIRE", args))
}

// ExpireAt https://redis.io/commands/expireat
// Command: EXPIREAT key timestamp [NX|XX|GT|LT]
// Integer reply: 1 if the timeout was set, 0 if the timeout was not set.
func (c *keyOps) ExpireAt(ctx context.Context, key string, expireAt int64, args ...any) (int64, error) {
	args = append([]any{key, expireAt}, args)
	return toInt64(c.driver.Exec(ctx, "EXPIREAT", args))
}

// Keys https://redis.io/commands/keys
// Command: KEYS pattern
// Array reply: list of keys matching pattern.
func (c *keyOps) Keys(ctx context.Context, pattern string) ([]string, error) {
	args := []any{pattern}
	return toStringSlice(c.driver.Exec(ctx, "KEYS", args))
}

// Persist https://redis.io/commands/persist
// Command: PERSIST key
// Integer reply: 1 if the timeout was removed,
// 0 if key does not exist or does not have an associated timeout.
func (c *keyOps) Persist(ctx context.Context, key string) (int64, error) {
	args := []any{key}
	return toInt64(c.driver.Exec(ctx, "PERSIST", args))
}

// PExpire https://redis.io/commands/pexpire
// Command: PEXPIRE key milliseconds [NX|XX|GT|LT]
// Integer reply: 1 if the timeout was set, 0 if the timeout was not set.
func (c *keyOps) PExpire(ctx context.Context, key string, expire int64, args ...any) (int64, error) {
	args = append([]any{key, expire}, args)
	return toInt64(c.driver.Exec(ctx, "PEXPIRE", args))
}

// PExpireAt https://redis.io/commands/pexpireat
// Command: PEXPIREAT key milliseconds-timestamp [NX|XX|GT|LT]
// Integer reply: 1 if the timeout was set, 0 if the timeout was not set.
func (c *keyOps) PExpireAt(ctx context.Context, key string, expireAt int64, args ...any) (int64, error) {
	args = append([]any{key, expireAt}, args)
	return toInt64(c.driver.Exec(ctx, "PEXPIREAT", args))
}

// PTTL https://redis.io/commands/pttl
// Command: PTTL key
// Integer reply: TTL in milliseconds, -1 if the key exists
// but has no associated expire, -2 if the key does not exist.
func (c *keyOps) PTTL(ctx context.Context, key string) (int64, error) {
	args := []any{key}
	return toInt64(c.driver.Exec(ctx, "PTTL", args))
}

// RandomKey https://redis.io/commands/randomkey
// Command: RANDOMKEY
// Bulk string reply: the random key, or nil when the database is empty.
func (c *keyOps) RandomKey(ctx context.Context) (string, error) {
	return toString(c.driver.Exec(ctx, "RANDOMKEY", nil))
}

// Rename https://redis.io/commands/rename
// Command: RENAME key newkey
// Simple string reply.
func (c *keyOps) Rename(ctx context.Context, key, newKey string) (string, error) {
	args := []any{key, newKey}
	return toString(c.driver.Exec(ctx, "RENAME", args))
}

// RenameNX https://redis.io/commands/renamenx
// Command: RENAMENX key newkey
// Integer reply: 1 if key was renamed to newKey, 0 if newKey already exists.
func (c *keyOps) RenameNX(ctx context.Context, key, newKey string) (int64, error) {
	args := []any{key, newKey}
	return toInt64(c.driver.Exec(ctx, "RENAMENX", args))
}

// Touch https://redis.io/commands/touch
// Command: TOUCH key [key ...]
// Integer reply: The number of keys that were touched.
func (c *keyOps) Touch(ctx context.Context, keys ...string) (int64, error) {
	var args []any
	for _, key := range keys {
		args = append(args, key)
	}
	return toInt64(c.driver.Exec(ctx, "TOUCH", args))
}

// TTL https://redis.io/commands/ttl
// Command: TTL key
// Integer reply: TTL in seconds, -1 if the key exists
// but has no associated expire, -2 if the key does not exist.
func (c *keyOps) TTL(ctx context.Context, key string) (int64, error) {
	args := []any{key}
	return toInt64(c.driver.Exec(ctx, "TTL", args))
}

// Type https://redis.io/commands/type
// Command: TYPE key
// Simple string reply: type of key, or none when key does not exist.
func (c *keyOps) Type(ctx context.Context, key string) (string, error) {
	args := []any{key}
	return toString(c.driver.Exec(ctx, "TYPE", args))
}
