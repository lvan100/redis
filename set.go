package redis

import "context"

type SetOps interface {

	// SAdd https://redis.io/commands/sadd
	// Command: SADD key member [member ...]
	// Integer reply: the number of elements that were added to the set,
	// not including all the elements already present in the set.
	SAdd(ctx context.Context, key string, members ...any) *Int64Replier

	// SCard https://redis.io/commands/scard
	// Command: SCARD key
	// Integer reply: the cardinality (number of elements) of the set,
	// or 0 if key does not exist.
	SCard(ctx context.Context, key string) *Int64Replier

	// SDiff https://redis.io/commands/sdiff
	// Command: SDIFF key [key ...]
	// Array reply: list with members of the resulting set.
	SDiff(ctx context.Context, keys ...string) *StringSliceReplier

	// SDiffStore https://redis.io/commands/sdiffstore
	// Command: SDIFFSTORE destination key [key ...]
	// Integer reply: the number of elements in the resulting set.
	SDiffStore(ctx context.Context, destination string, keys ...string) *Int64Replier

	// SInter https://redis.io/commands/sinter
	// Command: SINTER key [key ...]
	// Array reply: list with members of the resulting set.
	SInter(ctx context.Context, keys ...string) *StringSliceReplier

	// SInterStore https://redis.io/commands/sinterstore
	// Command: SINTERSTORE destination key [key ...]
	// Integer reply: the number of elements in the resulting set.
	SInterStore(ctx context.Context, destination string, keys ...string) *Int64Replier

	// SIsMember https://redis.io/commands/sismember
	// Command: SISMEMBER key member
	// Integer reply: 1 if the element is a member of the set,
	// 0 if the element is not a member of the set, or if key does not exist.
	SIsMember(ctx context.Context, key string, member any) *Int64Replier

	// SMembers https://redis.io/commands/smembers
	// Command: SMEMBERS key
	// Array reply: all elements of the set.
	SMembers(ctx context.Context, key string) *StringSliceReplier

	// SMIsMember https://redis.io/commands/smismember
	// Command: SMISMEMBER key member [member ...]
	// Array reply: list representing the membership of the given elements,
	// in the same order as they are requested.
	SMIsMember(ctx context.Context, key string, members ...any) *Int64SliceReplier

	// SMove https://redis.io/commands/smove
	// Command: SMOVE source destination member
	// Integer reply: 1 if the element is moved, 0 if the element
	// is not a member of source and no operation was performed.
	SMove(ctx context.Context, source, destination string, member any) *Int64Replier

	// SPop https://redis.io/commands/spop
	// Command: SPOP key [count]
	// Bulk string reply: the removed member, or nil when key does not exist.
	SPop(ctx context.Context, key string) *StringReplier

	// SPopN https://redis.io/commands/spop
	// Command: SPOP key [count]
	// Array reply: the removed members, or an empty array when key does not exist.
	SPopN(ctx context.Context, key string, count int64) *StringSliceReplier

	// SRandMember https://redis.io/commands/srandmember
	// Command: SRANDMEMBER key [count]
	// Returns a Bulk Reply with the randomly selected element,
	// or nil when key does not exist.
	SRandMember(ctx context.Context, key string) *StringReplier

	// SRandMemberN https://redis.io/commands/srandmember
	// Command: SRANDMEMBER key [count]
	// Returns an array of elements, or an empty array when key does not exist.
	SRandMemberN(ctx context.Context, key string, count int64) *StringSliceReplier

	// SRem https://redis.io/commands/srem
	// Command: SREM key member [member ...]
	// Integer reply: the number of members that were removed from the set,
	// not including non-existing members.
	SRem(ctx context.Context, key string, members ...any) *Int64Replier

	// SUnion https://redis.io/commands/sunion
	// Command: SUNION key [key ...]
	// Array reply: list with members of the resulting set.
	SUnion(ctx context.Context, keys ...string) *StringSliceReplier

	// SUnionStore https://redis.io/commands/sunionstore
	// Command: SUNIONSTORE destination key [key ...]
	// Integer reply: the number of elements in the resulting set.
	SUnionStore(ctx context.Context, destination string, keys ...string) *Int64Replier
}

/////////////////////////////

var _ SetOps = (*setOps)(nil)

type setOps struct {
	driver Driver
}

func (c *setOps) SAdd(ctx context.Context, key string, members ...any) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "SADD",
			args:   append([]any{key}, members...),
		},
	}
}

func (c *setOps) SCard(ctx context.Context, key string) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "SCARD",
			args:   []any{key},
		},
	}
}

func (c *setOps) SDiff(ctx context.Context, keys ...string) *StringSliceReplier {
	var args []any
	for _, key := range keys {
		args = append(args, key)
	}
	return &StringSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "SDIFF",
			args:   args,
		},
	}
}

func (c *setOps) SDiffStore(ctx context.Context, destination string, keys ...string) *Int64Replier {
	args := []any{destination}
	for _, key := range keys {
		args = append(args, key)
	}
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "SDIFFSTORE",
			args:   args,
		},
	}
}

func (c *setOps) SInter(ctx context.Context, keys ...string) *StringSliceReplier {
	var args []any
	for _, key := range keys {
		args = append(args, key)
	}
	return &StringSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "SINTER",
			args:   args,
		},
	}
}

func (c *setOps) SInterStore(ctx context.Context, destination string, keys ...string) *Int64Replier {
	args := []any{destination}
	for _, key := range keys {
		args = append(args, key)
	}
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "SINTERSTORE",
			args:   args,
		},
	}
}

func (c *setOps) SIsMember(ctx context.Context, key string, member any) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "SISMEMBER",
			args:   []any{key, member},
		},
	}
}

func (c *setOps) SMembers(ctx context.Context, key string) *StringSliceReplier {
	return &StringSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "SMEMBERS",
			args:   []any{key},
		},
	}
}

func (c *setOps) SMIsMember(ctx context.Context, key string, members ...any) *Int64SliceReplier {
	args := []any{key}
	for _, member := range members {
		args = append(args, member)
	}
	return &Int64SliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "SMISMEMBER",
			args:   args,
		},
	}
}

func (c *setOps) SMove(ctx context.Context, source, destination string, member any) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "SMOVE",
			args:   []any{source, destination, member},
		},
	}
}

func (c *setOps) SPop(ctx context.Context, key string) *StringReplier {
	return &StringReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "SPOP",
			args:   []any{key},
		},
	}
}

func (c *setOps) SPopN(ctx context.Context, key string, count int64) *StringSliceReplier {
	return &StringSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "SPOP",
			args:   []any{key, count},
		},
	}
}

func (c *setOps) SRandMember(ctx context.Context, key string) *StringReplier {
	return &StringReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "SRANDMEMBER",
			args:   []any{key},
		},
	}
}

func (c *setOps) SRandMemberN(ctx context.Context, key string, count int64) *StringSliceReplier {
	return &StringSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "SRANDMEMBER",
			args:   []any{key, count},
		},
	}
}

func (c *setOps) SRem(ctx context.Context, key string, members ...any) *Int64Replier {
	args := []any{key}
	for _, member := range members {
		args = append(args, member)
	}
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "SREM",
			args:   args,
		},
	}
}

func (c *setOps) SUnion(ctx context.Context, keys ...string) *StringSliceReplier {
	var args []any
	for _, key := range keys {
		args = append(args, key)
	}
	return &StringSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "SUNION",
			args:   args,
		},
	}
}

func (c *setOps) SUnionStore(ctx context.Context, destination string, keys ...string) *Int64Replier {
	args := []any{destination}
	for _, key := range keys {
		args = append(args, key)
	}
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "SUNIONSTORE",
			args:   args,
		},
	}
}
