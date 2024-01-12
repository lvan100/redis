package redis

import "context"

type BitmapOps interface {

	// BitCount https://redis.io/commands/bitcount
	// Command: BITCOUNT key [start end]
	// Integer reply: The number of bits set to 1.
	BitCount(ctx context.Context, key string, args ...any) (int64, error)

	// BitOpAnd https://redis.io/commands/bitop
	// Command: BITOP AND destkey srckey1 srckey2 srckey3 ... srckeyN
	// Integer reply: The size of the string stored in the destination key,
	// that is equal to the size of the longest input string.
	BitOpAnd(ctx context.Context, destKey string, keys ...string) (int64, error)

	// BitOpOr https://redis.io/commands/bitop
	// Command: BITOP OR destkey srckey1 srckey2 srckey3 ... srckeyN
	// Integer reply: The size of the string stored in the destination key,
	// that is equal to the size of the longest input string.
	BitOpOr(ctx context.Context, destKey string, keys ...string) (int64, error)

	// BitOpXor https://redis.io/commands/bitop
	// Command: BITOP XOR destkey srckey1 srckey2 srckey3 ... srckeyN
	// Integer reply: The size of the string stored in the destination key,
	// that is equal to the size of the longest input string.
	BitOpXor(ctx context.Context, destKey string, keys ...string) (int64, error)

	// BitOpNot https://redis.io/commands/bitop
	// Command: BITOP NOT destkey srckey
	// Integer reply: The size of the string stored in the destination key,
	// that is equal to the size of the longest input string.
	BitOpNot(ctx context.Context, destKey string, key string) (int64, error)

	// BitPos https://redis.io/commands/bitpos
	// Command: BITPOS key bit [start [end]]
	// Integer reply: The command returns the position of the first bit
	// set to 1 or 0 according to the request.
	BitPos(ctx context.Context, key string, bit int64, args ...any) (int64, error)

	// GetBit https://redis.io/commands/getbit
	// Command: GETBIT key offset
	// Integer reply: the bit value stored at offset.
	GetBit(ctx context.Context, key string, offset int64) (int64, error)

	// SetBit https://redis.io/commands/setbit
	// Command: SETBIT key offset value
	// Integer reply: the original bit value stored at offset.
	SetBit(ctx context.Context, key string, offset int64, value int) (int64, error)
}

/////////////////////////

var _ BitmapOps = (*bitmapOps)(nil)

type bitmapOps struct {
	driver Driver
}

// BitCount https://redis.io/commands/bitcount
// Command: BITCOUNT key [start end]
// Integer reply: The number of bits set to 1.
func (c *bitmapOps) BitCount(ctx context.Context, key string, args ...any) (int64, error) {
	args = append([]any{key}, args...)
	return toInt64(c.driver.Exec(ctx, "BITCOUNT", args))
}

// BitOpAnd https://redis.io/commands/bitop
// Command: BITOP AND destkey srckey1 srckey2 srckey3 ... srckeyN
// Integer reply: The size of the string stored in the destination key,
// that is equal to the size of the longest input string.
func (c *bitmapOps) BitOpAnd(ctx context.Context, destKey string, keys ...string) (int64, error) {
	args := []any{"AND", destKey}
	for _, key := range keys {
		args = append(args, key)
	}
	return toInt64(c.driver.Exec(ctx, "BITOP", args))
}

// BitOpOr https://redis.io/commands/bitop
// Command: BITOP OR destkey srckey1 srckey2 srckey3 ... srckeyN
// Integer reply: The size of the string stored in the destination key,
// that is equal to the size of the longest input string.
func (c *bitmapOps) BitOpOr(ctx context.Context, destKey string, keys ...string) (int64, error) {
	args := []any{"OR", destKey}
	for _, key := range keys {
		args = append(args, key)
	}
	return toInt64(c.driver.Exec(ctx, "BITOP", args))
}

// BitOpXor https://redis.io/commands/bitop
// Command: BITOP XOR destkey srckey1 srckey2 srckey3 ... srckeyN
// Integer reply: The size of the string stored in the destination key,
// that is equal to the size of the longest input string.
func (c *bitmapOps) BitOpXor(ctx context.Context, destKey string, keys ...string) (int64, error) {
	args := []any{"XOR", destKey}
	for _, key := range keys {
		args = append(args, key)
	}
	return toInt64(c.driver.Exec(ctx, "BITOP", args))
}

// BitOpNot https://redis.io/commands/bitop
// Command: BITOP NOT destkey srckey
// Integer reply: The size of the string stored in the destination key,
// that is equal to the size of the longest input string.
func (c *bitmapOps) BitOpNot(ctx context.Context, destKey string, key string) (int64, error) {
	args := []any{"NOT", destKey, key}
	return toInt64(c.driver.Exec(ctx, "BITOP", args))
}

// BitPos https://redis.io/commands/bitpos
// Command: BITPOS key bit [start [end]]
// Integer reply: The command returns the position of the first bit
// set to 1 or 0 according to the request.
func (c *bitmapOps) BitPos(ctx context.Context, key string, bit int64, args ...any) (int64, error) {
	args = append([]any{key, bit}, args...)
	return toInt64(c.driver.Exec(ctx, "BITPOS", args))
}

// GetBit https://redis.io/commands/getbit
// Command: GETBIT key offset
// Integer reply: the bit value stored at offset.
func (c *bitmapOps) GetBit(ctx context.Context, key string, offset int64) (int64, error) {
	args := []any{key, offset}
	return toInt64(c.driver.Exec(ctx, "GETBIT", args))
}

// SetBit https://redis.io/commands/setbit
// Command: SETBIT key offset value
// Integer reply: the original bit value stored at offset.
func (c *bitmapOps) SetBit(ctx context.Context, key string, offset int64, value int) (int64, error) {
	args := []any{key, offset, value}
	return toInt64(c.driver.Exec(ctx, "SETBIT", args))
}
