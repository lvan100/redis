package redis

import "context"

type CmdBitCount struct {
	optStart[*CmdBitCount]
	optEnd[*CmdBitCount]
	Int64Replier
}

type CmdBitPos struct {
	optStart[*CmdBitPos]
	optEnd[*CmdBitPos]
	Int64Replier
}

type BitmapOps interface {

	// BitCount https://redis.io/commands/bitcount
	// Command: BITCOUNT key [start end]
	// Integer reply: The number of bits set to 1.
	BitCount(ctx context.Context, key string) *CmdBitCount

	// BitAnd https://redis.io/commands/bitop
	// Command: BITOP AND destkey srckey1 srckey2 srckey3 ... srckeyN
	// Integer reply: The size of the string stored in the destination key,
	// that is equal to the size of the longest input string.
	BitAnd(ctx context.Context, destKey string, keys ...string) *Int64Replier

	// BitOr https://redis.io/commands/bitop
	// Command: BITOP OR destkey srckey1 srckey2 srckey3 ... srckeyN
	// Integer reply: The size of the string stored in the destination key,
	// that is equal to the size of the longest input string.
	BitOr(ctx context.Context, destKey string, keys ...string) *Int64Replier

	// BitXor https://redis.io/commands/bitop
	// Command: BITOP XOR destkey srckey1 srckey2 srckey3 ... srckeyN
	// Integer reply: The size of the string stored in the destination key,
	// that is equal to the size of the longest input string.
	BitXor(ctx context.Context, destKey string, keys ...string) *Int64Replier

	// BitNot https://redis.io/commands/bitop
	// Command: BITOP NOT destkey srckey
	// Integer reply: The size of the string stored in the destination key,
	// that is equal to the size of the longest input string.
	BitNot(ctx context.Context, destKey string, key string) *Int64Replier

	// BitPos https://redis.io/commands/bitpos
	// Command: BITPOS key bit [start [end]]
	// Integer reply: The command returns the position of the first bit
	// set to 1 or 0 according to the request.
	BitPos(ctx context.Context, key string, bit int64) *CmdBitPos

	// GetBit https://redis.io/commands/getbit
	// Command: GETBIT key offset
	// Integer reply: the bit value stored at offset.
	GetBit(ctx context.Context, key string, offset int64) *Int64Replier

	// SetBit https://redis.io/commands/setbit
	// Command: SETBIT key offset value
	// Integer reply: the original bit value stored at offset.
	SetBit(ctx context.Context, key string, offset int64, value int) *Int64Replier
}

/////////////////////////

var _ BitmapOps = (*bitmapOps)(nil)

type bitmapOps struct {
	driver Driver
}

func (c *bitmapOps) BitCount(ctx context.Context, key string) *CmdBitCount {
	return &CmdBitCount{
		Int64Replier: Int64Replier{
			command: command{
				driver: c.driver,
				ctx:    ctx,
				cmd:    "BITCOUNT",
				args:   []any{key},
			},
		},
	}
}

func (c *bitmapOps) BitAnd(ctx context.Context, destKey string, keys ...string) *Int64Replier {
	args := []any{"AND", destKey}
	for _, key := range keys {
		args = append(args, key)
	}
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "BITOP",
			args:   args,
		},
	}
}

func (c *bitmapOps) BitOr(ctx context.Context, destKey string, keys ...string) *Int64Replier {
	args := []any{"OR", destKey}
	for _, key := range keys {
		args = append(args, key)
	}
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "BITOP",
			args:   args,
		},
	}
}

func (c *bitmapOps) BitXor(ctx context.Context, destKey string, keys ...string) *Int64Replier {
	args := []any{"XOR", destKey}
	for _, key := range keys {
		args = append(args, key)
	}
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "BITOP",
			args:   args,
		},
	}
}

func (c *bitmapOps) BitNot(ctx context.Context, destKey string, key string) *Int64Replier {
	args := []any{"NOT", destKey, key}
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "BITOP",
			args:   args,
		},
	}
}

func (c *bitmapOps) BitPos(ctx context.Context, key string, bit int64) *CmdBitPos {
	return &CmdBitPos{
		Int64Replier: Int64Replier{
			command: command{
				driver: c.driver,
				ctx:    ctx,
				cmd:    "BITPOS",
				args:   []any{key, bit},
			},
		},
	}
}

func (c *bitmapOps) GetBit(ctx context.Context, key string, offset int64) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "GETBIT",
			args:   []any{key, offset},
		},
	}
}

func (c *bitmapOps) SetBit(ctx context.Context, key string, offset int64, value int) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "SETBIT",
			args:   []any{key, offset, value},
		},
	}
}
