package redis

import "context"

type HashOps interface {

	// HDel https://redis.io/commands/hdel
	// Command: HDEL key field [field ...]
	// Integer reply: the number of fields that were removed
	// from the hash, not including specified but non-existing fields.
	HDel(ctx context.Context, key string, fields ...string) (int64, error)

	// HExists https://redis.io/commands/hexists
	// Command: HEXISTS key field
	// Integer reply: 1 if the hash contains field,
	// 0 if the hash does not contain field, or key does not exist.
	HExists(ctx context.Context, key, field string) (int64, error)

	// HGet https://redis.io/commands/hget
	// Command: HGET key field
	// Bulk string reply: the value associated with field,
	// or nil when field is not present in the hash or key does not exist.
	HGet(ctx context.Context, key string, field string) (string, error)

	// HGetAll https://redis.io/commands/hgetall
	// Command: HGETALL key
	// Array reply: list of fields and their values stored
	// in the hash, or an empty list when key does not exist.
	HGetAll(ctx context.Context, key string) (map[string]string, error)

	// HIncrBy https://redis.io/commands/hincrby
	// Command: HINCRBY key field increment
	// Integer reply: the value at field after the increment operation.
	HIncrBy(ctx context.Context, key, field string, incr int64) (int64, error)

	// HIncrByFloat https://redis.io/commands/hincrbyfloat
	// Command: HINCRBYFLOAT key field increment
	// Bulk string reply: the value of field after the increment.
	HIncrByFloat(ctx context.Context, key, field string, incr float64) (float64, error)

	// HKeys https://redis.io/commands/hkeys
	// Command: HKEYS key
	// Array reply: list of fields in the hash, or an empty list when key does not exist.
	HKeys(ctx context.Context, key string) ([]string, error)

	// HLen https://redis.io/commands/hlen
	// Command: HLEN key
	// Integer reply: number of fields in the hash, or 0 when key does not exist.
	HLen(ctx context.Context, key string) (int64, error)

	// HMGet https://redis.io/commands/hmget
	// Command: HMGET key field [field ...]
	// Array reply: list of values associated with the
	// given fields, in the same order as they are requested.
	HMGet(ctx context.Context, key string, fields ...string) ([]any, error)

	// HSet https://redis.io/commands/hset
	// Command: HSET key field value [field value ...]
	// Integer reply: The number of fields that were added.
	HSet(ctx context.Context, key string, args ...any) (int64, error)

	// HSetNX https://redis.io/commands/hsetnx
	// Command: HSETNX key field value
	// Integer reply: 1 if field is a new field in the hash and value was set,
	// 0 if field already exists in the hash and no operation was performed.
	HSetNX(ctx context.Context, key, field string, value any) (int64, error)

	// HStrLen https://redis.io/commands/hstrlen
	// Command: HSTRLEN key field
	// Integer reply: the string length of the value associated with field,
	// or zero when field is not present in the hash or key does not exist at all.
	HStrLen(ctx context.Context, key, field string) (int64, error)

	// HVals https://redis.io/commands/hvals
	// Command: HVALS key
	// Array reply: list of values in the hash, or an empty list when key does not exist.
	HVals(ctx context.Context, key string) ([]string, error)
}

//////////////////////

var _ HashOps = (*hashOps)(nil)

type hashOps struct {
	driver Driver
}

// HDel https://redis.io/commands/hdel
// Command: HDEL key field [field ...]
// Integer reply: the number of fields that were removed
// from the hash, not including specified but non-existing fields.
func (c *hashOps) HDel(ctx context.Context, key string, fields ...string) (int64, error) {
	args := []any{key}
	for _, field := range fields {
		args = append(args, field)
	}
	return toInt64(c.driver.Exec(ctx, "HDEL", args))
}

// HExists https://redis.io/commands/hexists
// Command: HEXISTS key field
// Integer reply: 1 if the hash contains field,
// 0 if the hash does not contain field, or key does not exist.
func (c *hashOps) HExists(ctx context.Context, key, field string) (int64, error) {
	args := []any{key, field}
	return toInt64(c.driver.Exec(ctx, "HEXISTS", args))
}

// HGet https://redis.io/commands/hget
// Command: HGET key field
// Bulk string reply: the value associated with field,
// or nil when field is not present in the hash or key does not exist.
func (c *hashOps) HGet(ctx context.Context, key string, field string) (string, error) {
	args := []any{key, field}
	return toString(c.driver.Exec(ctx, "HGET", args))
}

// HGetAll https://redis.io/commands/hgetall
// Command: HGETALL key
// Array reply: list of fields and their values stored
// in the hash, or an empty list when key does not exist.
func (c *hashOps) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	args := []any{key}
	return toStringMap(c.driver.Exec(ctx, "HGETALL", args))
}

// HIncrBy https://redis.io/commands/hincrby
// Command: HINCRBY key field increment
// Integer reply: the value at field after the increment operation.
func (c *hashOps) HIncrBy(ctx context.Context, key, field string, incr int64) (int64, error) {
	args := []any{key, field, incr}
	return toInt64(c.driver.Exec(ctx, "HINCRBY", args))
}

// HIncrByFloat https://redis.io/commands/hincrbyfloat
// Command: HINCRBYFLOAT key field increment
// Bulk string reply: the value of field after the increment.
func (c *hashOps) HIncrByFloat(ctx context.Context, key, field string, incr float64) (float64, error) {
	args := []any{key, field, incr}
	return toFloat64(c.driver.Exec(ctx, "HINCRBYFLOAT", args))
}

// HKeys https://redis.io/commands/hkeys
// Command: HKEYS key
// Array reply: list of fields in the hash, or an empty list when key does not exist.
func (c *hashOps) HKeys(ctx context.Context, key string) ([]string, error) {
	args := []any{key}
	return toStringSlice(c.driver.Exec(ctx, "HKEYS", args))
}

// HLen https://redis.io/commands/hlen
// Command: HLEN key
// Integer reply: number of fields in the hash, or 0 when key does not exist.
func (c *hashOps) HLen(ctx context.Context, key string) (int64, error) {
	args := []any{key}
	return toInt64(c.driver.Exec(ctx, "HLEN", args))
}

// HMGet https://redis.io/commands/hmget
// Command: HMGET key field [field ...]
// Array reply: list of values associated with the
// given fields, in the same order as they are requested.
func (c *hashOps) HMGet(ctx context.Context, key string, fields ...string) ([]any, error) {
	args := []any{key}
	for _, field := range fields {
		args = append(args, field)
	}
	return toSlice(c.driver.Exec(ctx, "HMGET", args))
}

// HSet https://redis.io/commands/hset
// Command: HSET key field value [field value ...]
// Integer reply: The number of fields that were added.
func (c *hashOps) HSet(ctx context.Context, key string, args ...any) (int64, error) {
	args = append([]any{key}, args)
	return toInt64(c.driver.Exec(ctx, "HSET", args))
}

// HSetNX https://redis.io/commands/hsetnx
// Command: HSETNX key field value
// Integer reply: 1 if field is a new field in the hash and value was set,
// 0 if field already exists in the hash and no operation was performed.
func (c *hashOps) HSetNX(ctx context.Context, key, field string, value any) (int64, error) {
	args := []any{key, field, value}
	return toInt64(c.driver.Exec(ctx, "HSETNX", args))
}

// HStrLen https://redis.io/commands/hstrlen
// Command: HSTRLEN key field
// Integer reply: the string length of the value associated with field,
// or zero when field is not present in the hash or key does not exist at all.
func (c *hashOps) HStrLen(ctx context.Context, key, field string) (int64, error) {
	args := []any{key, field}
	return toInt64(c.driver.Exec(ctx, "HSTRLEN", args))
}

// HVals https://redis.io/commands/hvals
// Command: HVALS key
// Array reply: list of values in the hash, or an empty list when key does not exist.
func (c *hashOps) HVals(ctx context.Context, key string) ([]string, error) {
	args := []any{key}
	return toStringSlice(c.driver.Exec(ctx, "HVALS", args))
}
