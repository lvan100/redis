package redis

import "context"

type CmdLMove struct {
	optLeft[*CmdLMove]
	optRight[*CmdLMove]
	StringReplier
}

type CmdLPos struct {
	optRank[*CmdLPos]
	Int64Replier
}

type CmdLPosN struct {
	optRank[*CmdLPos]
	optCount[*CmdLPos]
	optMaxLen[*CmdLPos]
	Int64SliceReplier
}

type ListOps interface {

	// LIndex https://redis.io/commands/lindex
	// Command: LINDEX key index
	// Bulk string reply: the requested element, or nil when index is out of range.
	LIndex(ctx context.Context, key string, index int64) *StringReplier

	// LInsertBefore https://redis.io/commands/linsert
	// Command: LINSERT key BEFORE|AFTER pivot element
	// Integer reply: the length of the list after the
	// insert operation, or -1 when the value pivot was not found.
	LInsertBefore(ctx context.Context, key string, pivot, value any) *Int64Replier

	// LInsertAfter https://redis.io/commands/linsert
	// Command: LINSERT key BEFORE|AFTER pivot element
	// Integer reply: the length of the list after the
	// insert operation, or -1 when the value pivot was not found.
	LInsertAfter(ctx context.Context, key string, pivot, value any) *Int64Replier

	// LLen https://redis.io/commands/llen
	// Command: LLEN key
	// Integer reply: the length of the list at key.
	LLen(ctx context.Context, key string) *Int64Replier

	// LMove https://redis.io/commands/lmove
	// Command: LMOVE source destination LEFT|RIGHT LEFT|RIGHT
	// Bulk string reply: the element being popped and pushed.
	LMove(ctx context.Context, source, destination string) *CmdLMove

	// LPop https://redis.io/commands/lpop
	// Command: LPOP key [count]
	// Bulk string reply: the value of the first element, or nil when key does not exist.
	LPop(ctx context.Context, key string) *StringReplier

	// LPopN https://redis.io/commands/lpop
	// Command: LPOP key [count]
	// Array reply: list of popped elements, or nil when key does not exist.
	LPopN(ctx context.Context, key string, count int) *StringSliceReplier

	// LPos https://redis.io/commands/lpos
	// Command: LPOS key element [RANK rank] [COUNT num-matches] [MAXLEN len]
	// The command returns the integer representing the matching element,
	// or nil if there is no match. However, if the COUNT option is given
	// the command returns an array (empty if there are no matches).
	LPos(ctx context.Context, key string, value any) *CmdLPos

	// LPosN https://redis.io/commands/lpos
	// Command: LPOS key element [RANK rank] [COUNT num-matches] [MAXLEN len]
	// The command returns the integer representing the matching element,
	// or nil if there is no match. However, if the COUNT option is given
	// the command returns an array (empty if there are no matches).
	LPosN(ctx context.Context, key string, value any, count int64) *CmdLPosN

	// LPush https://redis.io/commands/lpush
	// Command: LPUSH key element [element ...]
	// Integer reply: the length of the list after the push operations.
	LPush(ctx context.Context, key string, values ...any) *Int64Replier

	// LPushX https://redis.io/commands/lpushx
	// Command: LPUSHX key element [element ...]
	// Integer reply: the length of the list after the push operation.
	LPushX(ctx context.Context, key string, values ...any) *Int64Replier

	// LRange https://redis.io/commands/lrange
	// Command: LRANGE key start stop
	// Array reply: list of elements in the specified range.
	LRange(ctx context.Context, key string, start, stop int64) *StringSliceReplier

	// LRem https://redis.io/commands/lrem
	// Command: LREM key count element
	// Integer reply: the number of removed elements.
	LRem(ctx context.Context, key string, count int64, value any) *Int64Replier

	// LSet https://redis.io/commands/lset
	// Command: LSET key index element
	// Simple string reply
	LSet(ctx context.Context, key string, index int64, value any) *StringReplier

	// LTrim https://redis.io/commands/ltrim
	// Command: LTRIM key start stop
	// Simple string reply
	LTrim(ctx context.Context, key string, start, stop int64) *StringReplier

	// RPop https://redis.io/commands/rpop
	// Command: RPOP key [count]
	// Bulk string reply: the value of the last element, or nil when key does not exist.
	RPop(ctx context.Context, key string) *StringReplier

	// RPopN https://redis.io/commands/rpop
	// Command: RPOP key [count]
	// Array reply: list of popped elements, or nil when key does not exist.
	RPopN(ctx context.Context, key string, count int) *StringSliceReplier

	// RPopLPush https://redis.io/commands/rpoplpush
	// Command: RPOPLPUSH source destination
	// Bulk string reply: the element being popped and pushed.
	RPopLPush(ctx context.Context, source, destination string) *StringReplier

	// RPush https://redis.io/commands/rpush
	// Command: RPUSH key element [element ...]
	// Integer reply: the length of the list after the push operation.
	RPush(ctx context.Context, key string, values ...any) *Int64Replier

	// RPushX https://redis.io/commands/rpushx
	// Command: RPUSHX key element [element ...]
	// Integer reply: the length of the list after the push operation.
	RPushX(ctx context.Context, key string, values ...any) *Int64Replier
}

//////////////////////////////

var _ ListOps = (*listOps)(nil)

type listOps struct {
	driver Driver
}

func (c *listOps) LIndex(ctx context.Context, key string, index int64) *StringReplier {
	return &StringReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "LINDEX",
			args:   []any{key, index},
		},
	}
}

func (c *listOps) LInsertBefore(ctx context.Context, key string, pivot, value any) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "LINSERT",
			args:   []any{key, "BEFORE", pivot, value},
		},
	}
}

func (c *listOps) LInsertAfter(ctx context.Context, key string, pivot, value any) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "LINSERT",
			args:   []any{key, "AFTER", pivot, value},
		},
	}
}

func (c *listOps) LLen(ctx context.Context, key string) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "LLEN",
			args:   []any{key},
		},
	}
}

func (c *listOps) LMove(ctx context.Context, source, destination string) *CmdLMove {
	return &CmdLMove{
		StringReplier: StringReplier{
			command: command{
				driver: c.driver,
				ctx:    ctx,
				cmd:    "LMOVE",
				args:   []any{source, destination},
			},
		},
	}
}

func (c *listOps) LPop(ctx context.Context, key string) *StringReplier {
	return &StringReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "LPOP",
			args:   []any{key},
		},
	}
}

func (c *listOps) LPopN(ctx context.Context, key string, count int) *StringSliceReplier {
	return &StringSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "LPOP",
			args:   []any{key, count},
		},
	}
}

func (c *listOps) LPos(ctx context.Context, key string, value any) *CmdLPos {
	return &CmdLPos{
		Int64Replier: Int64Replier{
			command: command{
				driver: c.driver,
				ctx:    ctx,
				cmd:    "LPOP",
				args:   []any{key, value},
			},
		},
	}
}

func (c *listOps) LPosN(ctx context.Context, key string, value any, count int64) *CmdLPosN {
	return &CmdLPosN{
		Int64SliceReplier: Int64SliceReplier{
			command: command{
				driver: c.driver,
				ctx:    ctx,
				cmd:    "LPOS",
				args:   []any{key, value, "COUNT", count},
			},
		},
	}
}

func (c *listOps) LPush(ctx context.Context, key string, values ...any) *Int64Replier {
	args := []any{key}
	for _, value := range values {
		args = append(args, value)
	}
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "LPUSH",
			args:   args,
		},
	}
}

func (c *listOps) LPushX(ctx context.Context, key string, values ...any) *Int64Replier {
	args := []any{key}
	for _, value := range values {
		args = append(args, value)
	}
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "LPUSHX",
			args:   args,
		},
	}
}

func (c *listOps) LRange(ctx context.Context, key string, start, stop int64) *StringSliceReplier {
	return &StringSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "LRANGE",
			args:   []any{key, start, stop},
		},
	}
}

func (c *listOps) LRem(ctx context.Context, key string, count int64, value any) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "LREM",
			args:   []any{key, count, value},
		},
	}
}

func (c *listOps) LSet(ctx context.Context, key string, index int64, value any) *StringReplier {
	return &StringReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "LSET",
			args:   []any{key, index, value},
		},
	}
}

func (c *listOps) LTrim(ctx context.Context, key string, start, stop int64) *StringReplier {
	return &StringReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "LTRIM",
			args:   []any{key, start, stop},
		},
	}
}

func (c *listOps) RPop(ctx context.Context, key string) *StringReplier {
	return &StringReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "RPOP",
			args:   []any{key},
		},
	}
}

func (c *listOps) RPopN(ctx context.Context, key string, count int) *StringSliceReplier {
	return &StringSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "RPOP",
			args:   []any{key, count},
		},
	}
}

func (c *listOps) RPopLPush(ctx context.Context, source, destination string) *StringReplier {
	return &StringReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "RPOPLPUSH",
			args:   []any{source, destination},
		},
	}
}

func (c *listOps) RPush(ctx context.Context, key string, values ...any) *Int64Replier {
	args := []any{key}
	for _, value := range values {
		args = append(args, value)
	}
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "RPUSH",
			args:   args,
		},
	}
}

func (c *listOps) RPushX(ctx context.Context, key string, values ...any) *Int64Replier {
	args := []any{key}
	for _, value := range values {
		args = append(args, value)
	}
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "RPUSHX",
			args:   args,
		},
	}
}
