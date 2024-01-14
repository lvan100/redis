package redis

import "context"

type ZItem struct {
	Member any
	Score  float64
}

type SortedSetOps interface {

	// ZAdd https://redis.io/commands/zadd
	// Command: ZADD key [NX|XX] [GT|LT] [CH] [INCR] score member [score member ...]
	// Integer reply, the number of elements added to the
	// sorted set (excluding score updates).
	ZAdd(ctx context.Context, key string, args ...any) *Int64Replier

	// ZCard https://redis.io/commands/zcard
	// Command: ZCARD key
	// Integer reply: the cardinality (number of elements)
	// of the sorted set, or 0 if key does not exist.
	ZCard(ctx context.Context, key string) *Int64Replier

	// ZCount https://redis.io/commands/zcount
	// Command: ZCOUNT key min max
	// Integer reply: the number of elements in the specified score range.
	ZCount(ctx context.Context, key, min, max string) *Int64Replier

	// ZDiff https://redis.io/commands/zdiff
	// Command: ZDIFF numkeys key [key ...] [WITHSCORES]
	// Array reply: the result of the difference.
	ZDiff(ctx context.Context, keys ...string) *StringSliceReplier

	// ZDiffWithScores https://redis.io/commands/zdiff
	// Command: ZDIFF numkeys key [key ...] [WITHSCORES]
	// Array reply: the result of the difference.
	ZDiffWithScores(ctx context.Context, keys ...string) *ZItemSliceReplier

	// ZIncrBy https://redis.io/commands/zincrby
	// Command: ZINCRBY key increment member
	// Bulk string reply: the new score of member
	// (a double precision floating point number), represented as string.
	ZIncrBy(ctx context.Context, key string, increment float64, member string) *Float64Replier

	// ZInter https://redis.io/commands/zinter
	// Command: ZINTER numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX] [WITHSCORES]
	// Array reply: the result of intersection.
	ZInter(ctx context.Context, args ...any) *StringSliceReplier

	// ZInterWithScores https://redis.io/commands/zinter
	// Command: ZINTER numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX] [WITHSCORES]
	// Array reply: the result of intersection.
	ZInterWithScores(ctx context.Context, args ...any) *ZItemSliceReplier

	// ZLexCount https://redis.io/commands/zlexcount
	// Command: ZLEXCOUNT key min max
	// Integer reply: the number of elements in the specified score range.
	ZLexCount(ctx context.Context, key, min, max string) *Int64Replier

	// ZMScore https://redis.io/commands/zmscore
	// Command: ZMSCORE key member [member ...]
	// Array reply: list of scores or nil associated with the specified member
	// values (a double precision floating point number), represented as strings.
	ZMScore(ctx context.Context, key string, members ...string) *Float64SliceReplier

	// ZPopMax https://redis.io/commands/zpopmax
	// Command: ZPOPMAX key [count]
	// Array reply: list of popped elements and scores.
	ZPopMax(ctx context.Context, key string) *ZItemSliceReplier

	// ZPopMaxN https://redis.io/commands/zpopmax
	// Command: ZPOPMAX key [count]
	// Array reply: list of popped elements and scores.
	ZPopMaxN(ctx context.Context, key string, count int64) *ZItemSliceReplier

	// ZPopMin https://redis.io/commands/zpopmin
	// Command: ZPOPMIN key [count]
	// Array reply: list of popped elements and scores.
	ZPopMin(ctx context.Context, key string) *ZItemSliceReplier

	// ZPopMinN https://redis.io/commands/zpopmin
	// Command: ZPOPMIN key [count]
	// Array reply: list of popped elements and scores.
	ZPopMinN(ctx context.Context, key string, count int64) *ZItemSliceReplier

	// ZRandMember https://redis.io/commands/zrandmember
	// Command: ZRANDMEMBER key [count [WITHSCORES]]
	// Bulk Reply with the randomly selected element, or nil when key does not exist.
	ZRandMember(ctx context.Context, key string) *StringReplier

	// ZRandMemberN https://redis.io/commands/zrandmember
	// Command: ZRANDMEMBER key [count [WITHSCORES]]
	// Bulk Reply with the randomly selected element, or nil when key does not exist.
	ZRandMemberN(ctx context.Context, key string, count int) *StringSliceReplier

	// ZRandMemberWithScores https://redis.io/commands/zrandmember
	// Command: ZRANDMEMBER key [count [WITHSCORES]]
	// Bulk Reply with the randomly selected element, or nil when key does not exist.
	ZRandMemberWithScores(ctx context.Context, key string, count int) *ZItemSliceReplier

	// ZRange https://redis.io/commands/zrange
	// Command: ZRANGE key min max [BYSCORE|BYLEX] [REV] [LIMIT offset count] [WITHSCORES]
	// Array reply: list of elements in the specified range.
	ZRange(ctx context.Context, key string, start, stop int64, args ...any) *StringSliceReplier

	// ZRangeWithScores https://redis.io/commands/zrange
	// Command: ZRANGE key min max [BYSCORE|BYLEX] [REV] [LIMIT offset count] [WITHSCORES]
	// Array reply: list of elements in the specified range.
	ZRangeWithScores(ctx context.Context, key string, start, stop int64, args ...any) *ZItemSliceReplier

	// ZRangeByLex https://redis.io/commands/zrangebylex
	// Command: ZRANGEBYLEX key min max [LIMIT offset count]
	// Array reply: list of elements in the specified score range.
	ZRangeByLex(ctx context.Context, key string, min, max string, args ...any) *StringSliceReplier

	// ZRangeByScore https://redis.io/commands/zrangebyscore
	// Command: ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT offset count]
	// Array reply: list of elements in the specified score range.
	ZRangeByScore(ctx context.Context, key string, min, max string, args ...any) *StringSliceReplier

	// ZRank https://redis.io/commands/zrank
	// Command: ZRANK key member
	// If member exists in the sorted set, Integer reply: the rank of member.
	// If member does not exist in the sorted set or key does not exist, Bulk string reply: nil.
	ZRank(ctx context.Context, key, member string) *Int64Replier

	// ZRem https://redis.io/commands/zrem
	// Command: ZREM key member [member ...]
	// Integer reply, The number of members removed from the sorted set, not including non existing members.
	ZRem(ctx context.Context, key string, members ...any) *Int64Replier

	// ZRemRangeByLex https://redis.io/commands/zremrangebylex
	// Command: ZREMRANGEBYLEX key min max
	// Integer reply: the number of elements removed.
	ZRemRangeByLex(ctx context.Context, key, min, max string) *Int64Replier

	// ZRemRangeByRank https://redis.io/commands/zremrangebyrank
	// Command: ZREMRANGEBYRANK key start stop
	// Integer reply: the number of elements removed.
	ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *Int64Replier

	// ZRemRangeByScore https://redis.io/commands/zremrangebyscore
	// Command: ZREMRANGEBYSCORE key min max
	// Integer reply: the number of elements removed.
	ZRemRangeByScore(ctx context.Context, key, min, max string) *Int64Replier

	// ZRevRange https://redis.io/commands/zrevrange
	// Command: ZREVRANGE key start stop [WITHSCORES]
	// Array reply: list of elements in the specified range.
	ZRevRange(ctx context.Context, key string, start, stop int64) *StringSliceReplier

	// ZRevRangeWithScores https://redis.io/commands/zrevrange
	// Command: ZREVRANGE key start stop [WITHSCORES]
	// Array reply: list of elements in the specified range.
	ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *StringSliceReplier

	// ZRevRangeByLex https://redis.io/commands/zrevrangebylex
	// Command: ZREVRANGEBYLEX key max min [LIMIT offset count]
	// Array reply: list of elements in the specified score range.
	ZRevRangeByLex(ctx context.Context, key string, min, max string, args ...any) *StringSliceReplier

	// ZRevRangeByScore https://redis.io/commands/zrevrangebyscore
	// Command: ZREVRANGEBYSCORE key max min [WITHSCORES] [LIMIT offset count]
	// Array reply: list of elements in the specified score range.
	ZRevRangeByScore(ctx context.Context, key string, min, max string, args ...any) *StringSliceReplier

	// ZRevRank https://redis.io/commands/zrevrank
	// Command: ZREVRANK key member
	// If member exists in the sorted set, Integer reply: the rank of member.
	// If member does not exist in the sorted set or key does not exist, Bulk string reply: nil.
	ZRevRank(ctx context.Context, key, member string) *Int64Replier

	// ZScore https://redis.io/commands/zscore
	// Command: ZSCORE key member
	// Bulk string reply: the score of member (a double precision floating point number), represented as string.
	ZScore(ctx context.Context, key, member string) *Float64Replier

	// ZUnion https://redis.io/commands/zunion
	// Command: ZUNION numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX] [WITHSCORES]
	// Array reply: the result of union.
	ZUnion(ctx context.Context, args ...any) *StringSliceReplier

	// ZUnionWithScores https://redis.io/commands/zunion
	// Command: ZUNION numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX] [WITHSCORES]
	// Array reply: the result of union.
	ZUnionWithScores(ctx context.Context, args ...any) *ZItemSliceReplier

	// ZUnionStore https://redis.io/commands/zunionstore
	// Command: ZUNIONSTORE destination numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX]
	// Integer reply: the number of elements in the resulting sorted set at destination.
	ZUnionStore(ctx context.Context, dest string, args ...any) *Int64Replier
}

type sortedSetOps struct {
	driver Driver
}

func (c *sortedSetOps) ZAdd(ctx context.Context, key string, args ...any) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZADD",
			args:   append([]any{key}, args...),
		},
	}
}

func (c *sortedSetOps) ZCard(ctx context.Context, key string) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZCARD",
			args:   []any{key},
		},
	}
}

func (c *sortedSetOps) ZCount(ctx context.Context, key, min, max string) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZCOUNT",
			args:   []any{key, min, max},
		},
	}
}

func (c *sortedSetOps) ZDiff(ctx context.Context, keys ...string) *StringSliceReplier {
	args := []any{len(keys)}
	for _, key := range keys {
		args = append(args, key)
	}
	return &StringSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZDIFF",
			args:   args,
		},
	}
}

func (c *sortedSetOps) ZDiffWithScores(ctx context.Context, keys ...string) *ZItemSliceReplier {
	args := []any{len(keys)}
	for _, key := range keys {
		args = append(args, key)
	}
	args = append(args, "WITHSCORES")
	return &ZItemSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZDIFF",
			args:   args,
		},
	}
}

func (c *sortedSetOps) ZIncrBy(ctx context.Context, key string, increment float64, member string) *Float64Replier {
	return &Float64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZINCRBY",
			args:   []any{key, increment, member},
		},
	}
}

func (c *sortedSetOps) ZInter(ctx context.Context, args ...any) *StringSliceReplier {
	return &StringSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZINTER",
			args:   append([]any{}, args...),
		},
	}
}

func (c *sortedSetOps) ZInterWithScores(ctx context.Context, args ...any) *ZItemSliceReplier {
	args = append([]any{}, args...)
	args = append(args, "WITHSCORES")
	return &ZItemSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZINTER",
			args:   args,
		},
	}
}

func (c *sortedSetOps) ZLexCount(ctx context.Context, key, min, max string) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZLEXCOUNT",
			args:   []any{key, min, max},
		},
	}
}

func (c *sortedSetOps) ZMScore(ctx context.Context, key string, members ...string) *Float64SliceReplier {
	args := []any{key}
	for _, member := range members {
		args = append(args, member)
	}
	return &Float64SliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZMSCORE",
			args:   args,
		},
	}
}

func (c *sortedSetOps) ZPopMax(ctx context.Context, key string) *ZItemSliceReplier {
	return &ZItemSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZPOPMAX",
			args:   []any{key},
		},
	}
}

func (c *sortedSetOps) ZPopMaxN(ctx context.Context, key string, count int64) *ZItemSliceReplier {
	return &ZItemSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZPOPMAX",
			args:   []any{key, count},
		},
	}
}

func (c *sortedSetOps) ZPopMin(ctx context.Context, key string) *ZItemSliceReplier {
	return &ZItemSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZPOPMIN",
			args:   []any{key},
		},
	}
}

func (c *sortedSetOps) ZPopMinN(ctx context.Context, key string, count int64) *ZItemSliceReplier {
	return &ZItemSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZPOPMIN",
			args:   []any{key, count},
		},
	}
}

func (c *sortedSetOps) ZRandMember(ctx context.Context, key string) *StringReplier {
	return &StringReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZRANDMEMBER",
			args:   []any{key},
		},
	}
}

func (c *sortedSetOps) ZRandMemberN(ctx context.Context, key string, count int) *StringSliceReplier {
	return &StringSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZRANDMEMBER",
			args:   []any{key, count},
		},
	}
}

func (c *sortedSetOps) ZRandMemberWithScores(ctx context.Context, key string, count int) *ZItemSliceReplier {
	return &ZItemSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZRANDMEMBER",
			args:   []any{key, count, "WITHSCORES"},
		},
	}
}

func (c *sortedSetOps) ZRange(ctx context.Context, key string, start, stop int64, args ...any) *StringSliceReplier {
	return &StringSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZRANGE",
			args:   append([]any{key, start, stop}, args...),
		},
	}
}

func (c *sortedSetOps) ZRangeWithScores(ctx context.Context, key string, start, stop int64, args ...any) *ZItemSliceReplier {
	args = append([]any{key, start, stop}, args...)
	args = append(args, "WITHSCORES")
	return &ZItemSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZRANGE",
			args:   args,
		},
	}
}

func (c *sortedSetOps) ZRangeByLex(ctx context.Context, key string, min, max string, args ...any) *StringSliceReplier {
	return &StringSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZRANGEBYLEX",
			args:   append([]any{key, min, max}, args...),
		},
	}
}

func (c *sortedSetOps) ZRangeByScore(ctx context.Context, key string, min, max string, args ...any) *StringSliceReplier {
	return &StringSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZRANGEBYSCORE",
			args:   append([]any{key, min, max}, args...),
		},
	}
}

func (c *sortedSetOps) ZRank(ctx context.Context, key, member string) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZRANK",
			args:   []any{key, member},
		},
	}
}

func (c *sortedSetOps) ZRem(ctx context.Context, key string, members ...any) *Int64Replier {
	args := []any{key}
	for _, member := range members {
		args = append(args, member)
	}
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZREM",
			args:   args,
		},
	}
}

func (c *sortedSetOps) ZRemRangeByLex(ctx context.Context, key, min, max string) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZREMRANGEBYLEX",
			args:   []any{key, min, max},
		},
	}
}

func (c *sortedSetOps) ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZREMRANGEBYRANK",
			args:   []any{key, start, stop},
		},
	}
}

func (c *sortedSetOps) ZRemRangeByScore(ctx context.Context, key, min, max string) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZREMRANGEBYSCORE",
			args:   []any{key, min, max},
		},
	}
}

func (c *sortedSetOps) ZRevRange(ctx context.Context, key string, start, stop int64) *StringSliceReplier {
	return &StringSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZREVRANGE",
			args:   []any{key, start, stop},
		},
	}
}

func (c *sortedSetOps) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *StringSliceReplier {
	return &StringSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZREVRANGE",
			args:   []any{key, start, stop, "WITHSCORES"},
		},
	}
}

func (c *sortedSetOps) ZRevRangeByLex(ctx context.Context, key string, min, max string, args ...any) *StringSliceReplier {
	return &StringSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZREVRANGEBYLEX",
			args:   append([]any{key, min, max}, args...),
		},
	}
}

func (c *sortedSetOps) ZRevRangeByScore(ctx context.Context, key string, min, max string, args ...any) *StringSliceReplier {
	return &StringSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZREVRANGEBYSCORE",
			args:   append([]any{key, min, max}, args...),
		},
	}
}

func (c *sortedSetOps) ZRevRank(ctx context.Context, key, member string) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZREVRANK",
			args:   []any{key, member},
		},
	}
}

func (c *sortedSetOps) ZScore(ctx context.Context, key, member string) *Float64Replier {
	return &Float64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZSCORE",
			args:   []any{key, member},
		},
	}
}

func (c *sortedSetOps) ZUnion(ctx context.Context, args ...any) *StringSliceReplier {
	return &StringSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZUNION",
			args:   args,
		},
	}
}

func (c *sortedSetOps) ZUnionWithScores(ctx context.Context, args ...any) *ZItemSliceReplier {
	return &ZItemSliceReplier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZUNION",
			args:   append(args, "WITHSCORES"),
		},
	}
}

func (c *sortedSetOps) ZUnionStore(ctx context.Context, dest string, args ...any) *Int64Replier {
	return &Int64Replier{
		command: command{
			driver: c.driver,
			ctx:    ctx,
			cmd:    "ZUNIONSTORE",
			args:   append([]any{dest}, args...),
		},
	}
}
