package redis

import "context"

type ZItem struct {
	Member any
	Score  float64
}

type ZSetOps interface {

	// ZAdd https://redis.io/commands/zadd
	// Command: ZADD key [NX|XX] [GT|LT] [CH] [INCR] score member [score member ...]
	// Integer reply, the number of elements added to the
	// sorted set (excluding score updates).
	ZAdd(ctx context.Context, key string, args ...any) (int64, error)

	// ZCard https://redis.io/commands/zcard
	// Command: ZCARD key
	// Integer reply: the cardinality (number of elements)
	// of the sorted set, or 0 if key does not exist.
	ZCard(ctx context.Context, key string) (int64, error)

	// ZCount https://redis.io/commands/zcount
	// Command: ZCOUNT key min max
	// Integer reply: the number of elements in the specified score range.
	ZCount(ctx context.Context, key, min, max string) (int64, error)

	// ZDiff https://redis.io/commands/zdiff
	// Command: ZDIFF numkeys key [key ...] [WITHSCORES]
	// Array reply: the result of the difference.
	ZDiff(ctx context.Context, keys ...string) ([]string, error)

	// ZDiffWithScores https://redis.io/commands/zdiff
	// Command: ZDIFF numkeys key [key ...] [WITHSCORES]
	// Array reply: the result of the difference.
	ZDiffWithScores(ctx context.Context, keys ...string) ([]ZItem, error)

	// ZIncrBy https://redis.io/commands/zincrby
	// Command: ZINCRBY key increment member
	// Bulk string reply: the new score of member
	// (a double precision floating point number), represented as string.
	ZIncrBy(ctx context.Context, key string, increment float64, member string) (float64, error)

	// ZInter https://redis.io/commands/zinter
	// Command: ZINTER numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX] [WITHSCORES]
	// Array reply: the result of intersection.
	ZInter(ctx context.Context, args ...any) ([]string, error)

	// ZInterWithScores https://redis.io/commands/zinter
	// Command: ZINTER numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX] [WITHSCORES]
	// Array reply: the result of intersection.
	ZInterWithScores(ctx context.Context, args ...any) ([]ZItem, error)

	// ZLexCount https://redis.io/commands/zlexcount
	// Command: ZLEXCOUNT key min max
	// Integer reply: the number of elements in the specified score range.
	ZLexCount(ctx context.Context, key, min, max string) (int64, error)

	// ZMScore https://redis.io/commands/zmscore
	// Command: ZMSCORE key member [member ...]
	// Array reply: list of scores or nil associated with the specified member
	// values (a double precision floating point number), represented as strings.
	ZMScore(ctx context.Context, key string, members ...string) ([]float64, error)

	// ZPopMax https://redis.io/commands/zpopmax
	// Command: ZPOPMAX key [count]
	// Array reply: list of popped elements and scores.
	ZPopMax(ctx context.Context, key string) ([]ZItem, error)

	// ZPopMaxN https://redis.io/commands/zpopmax
	// Command: ZPOPMAX key [count]
	// Array reply: list of popped elements and scores.
	ZPopMaxN(ctx context.Context, key string, count int64) ([]ZItem, error)

	// ZPopMin https://redis.io/commands/zpopmin
	// Command: ZPOPMIN key [count]
	// Array reply: list of popped elements and scores.
	ZPopMin(ctx context.Context, key string) ([]ZItem, error)

	// ZPopMinN https://redis.io/commands/zpopmin
	// Command: ZPOPMIN key [count]
	// Array reply: list of popped elements and scores.
	ZPopMinN(ctx context.Context, key string, count int64) ([]ZItem, error)

	// ZRandMember https://redis.io/commands/zrandmember
	// Command: ZRANDMEMBER key [count [WITHSCORES]]
	// Bulk Reply with the randomly selected element, or nil when key does not exist.
	ZRandMember(ctx context.Context, key string) (string, error)

	// ZRandMemberN https://redis.io/commands/zrandmember
	// Command: ZRANDMEMBER key [count [WITHSCORES]]
	// Bulk Reply with the randomly selected element, or nil when key does not exist.
	ZRandMemberN(ctx context.Context, key string, count int) ([]string, error)

	// ZRandMemberWithScores https://redis.io/commands/zrandmember
	// Command: ZRANDMEMBER key [count [WITHSCORES]]
	// Bulk Reply with the randomly selected element, or nil when key does not exist.
	ZRandMemberWithScores(ctx context.Context, key string, count int) ([]ZItem, error)

	// ZRange https://redis.io/commands/zrange
	// Command: ZRANGE key min max [BYSCORE|BYLEX] [REV] [LIMIT offset count] [WITHSCORES]
	// Array reply: list of elements in the specified range.
	ZRange(ctx context.Context, key string, start, stop int64, args ...any) ([]string, error)

	// ZRangeWithScores https://redis.io/commands/zrange
	// Command: ZRANGE key min max [BYSCORE|BYLEX] [REV] [LIMIT offset count] [WITHSCORES]
	// Array reply: list of elements in the specified range.
	ZRangeWithScores(ctx context.Context, key string, start, stop int64, args ...any) ([]ZItem, error)

	// ZRangeByLex https://redis.io/commands/zrangebylex
	// Command: ZRANGEBYLEX key min max [LIMIT offset count]
	// Array reply: list of elements in the specified score range.
	ZRangeByLex(ctx context.Context, key string, min, max string, args ...any) ([]string, error)

	// ZRangeByScore https://redis.io/commands/zrangebyscore
	// Command: ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT offset count]
	// Array reply: list of elements in the specified score range.
	ZRangeByScore(ctx context.Context, key string, min, max string, args ...any) ([]string, error)

	// ZRank https://redis.io/commands/zrank
	// Command: ZRANK key member
	// If member exists in the sorted set, Integer reply: the rank of member.
	// If member does not exist in the sorted set or key does not exist, Bulk string reply: nil.
	ZRank(ctx context.Context, key, member string) (int64, error)

	// ZRem https://redis.io/commands/zrem
	// Command: ZREM key member [member ...]
	// Integer reply, The number of members removed from the sorted set, not including non existing members.
	ZRem(ctx context.Context, key string, members ...any) (int64, error)

	// ZRemRangeByLex https://redis.io/commands/zremrangebylex
	// Command: ZREMRANGEBYLEX key min max
	// Integer reply: the number of elements removed.
	ZRemRangeByLex(ctx context.Context, key, min, max string) (int64, error)

	// ZRemRangeByRank https://redis.io/commands/zremrangebyrank
	// Command: ZREMRANGEBYRANK key start stop
	// Integer reply: the number of elements removed.
	ZRemRangeByRank(ctx context.Context, key string, start, stop int64) (int64, error)

	// ZRemRangeByScore https://redis.io/commands/zremrangebyscore
	// Command: ZREMRANGEBYSCORE key min max
	// Integer reply: the number of elements removed.
	ZRemRangeByScore(ctx context.Context, key, min, max string) (int64, error)

	// ZRevRange https://redis.io/commands/zrevrange
	// Command: ZREVRANGE key start stop [WITHSCORES]
	// Array reply: list of elements in the specified range.
	ZRevRange(ctx context.Context, key string, start, stop int64) ([]string, error)

	// ZRevRangeWithScores https://redis.io/commands/zrevrange
	// Command: ZREVRANGE key start stop [WITHSCORES]
	// Array reply: list of elements in the specified range.
	ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) ([]string, error)

	// ZRevRangeByLex https://redis.io/commands/zrevrangebylex
	// Command: ZREVRANGEBYLEX key max min [LIMIT offset count]
	// Array reply: list of elements in the specified score range.
	ZRevRangeByLex(ctx context.Context, key string, min, max string, args ...any) ([]string, error)

	// ZRevRangeByScore https://redis.io/commands/zrevrangebyscore
	// Command: ZREVRANGEBYSCORE key max min [WITHSCORES] [LIMIT offset count]
	// Array reply: list of elements in the specified score range.
	ZRevRangeByScore(ctx context.Context, key string, min, max string, args ...any) ([]string, error)

	// ZRevRank https://redis.io/commands/zrevrank
	// Command: ZREVRANK key member
	// If member exists in the sorted set, Integer reply: the rank of member.
	// If member does not exist in the sorted set or key does not exist, Bulk string reply: nil.
	ZRevRank(ctx context.Context, key, member string) (int64, error)

	// ZScore https://redis.io/commands/zscore
	// Command: ZSCORE key member
	// Bulk string reply: the score of member (a double precision floating point number), represented as string.
	ZScore(ctx context.Context, key, member string) (float64, error)

	// ZUnion https://redis.io/commands/zunion
	// Command: ZUNION numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX] [WITHSCORES]
	// Array reply: the result of union.
	ZUnion(ctx context.Context, args ...any) ([]string, error)

	// ZUnionWithScores https://redis.io/commands/zunion
	// Command: ZUNION numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX] [WITHSCORES]
	// Array reply: the result of union.
	ZUnionWithScores(ctx context.Context, args ...any) ([]ZItem, error)

	// ZUnionStore https://redis.io/commands/zunionstore
	// Command: ZUNIONSTORE destination numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX]
	// Integer reply: the number of elements in the resulting sorted set at destination.
	ZUnionStore(ctx context.Context, dest string, args ...any) (int64, error)
}

////////////////////////////////////////////////////
//
//var _ ZSetOps = (*orderedSetOps)(nil)
//
//type orderedSetOps struct {
//	driver Driver
//}
//
//// ZAdd https://redis.io/commands/zadd
//// Command: ZADD key [NX|XX] [GT|LT] [CH] [INCR] score member [score member ...]
//// Integer reply, the number of elements added to the
//// sorted set (excluding score updates).
//func (c *orderedSetOps) ZAdd(ctx context.Context, key string, args ...any) (int64, error) {
//	args = append([]any{"ZADD", key}, args...)
//	return toInt64(c.driver.Exec(ctx, args...)
//}
//
//// ZCard https://redis.io/commands/zcard
//// Command: ZCARD key
//// Integer reply: the cardinality (number of elements)
//// of the sorted set, or 0 if key does not exist.
//func (c *orderedSetOps) ZCard(ctx context.Context, key string) (int64, error) {
//	args := []any{"ZCARD", key}
//	return toInt64(c.driver.Exec(ctx, args...)
//}
//
//// ZCount https://redis.io/commands/zcount
//// Command: ZCOUNT key min max
//// Integer reply: the number of elements in the specified score range.
//func (c *orderedSetOps) ZCount(ctx context.Context, key, min, max string) (int64, error) {
//	args := []any{"ZCOUNT", key, min, max}
//	return toInt64(c.driver.Exec(ctx, args...)
//}
//
//// ZDiff https://redis.io/commands/zdiff
//// Command: ZDIFF numkeys key [key ...] [WITHSCORES]
//// Array reply: the result of the difference.
//func (c *orderedSetOps) ZDiff(ctx context.Context, keys ...string) ([]string, error) {
//	args := []any{"ZDIFF", len(keys)}
//	for _, key := range keys {
//		args = append(args, key)
//	}
//	return toString(c.driver.Exec(Slice(ctx, args...)
//}
//
//// ZDiffWithScores https://redis.io/commands/zdiff
//// Command: ZDIFF numkeys key [key ...] [WITHSCORES]
//// Array reply: the result of the difference.
//func (c *orderedSetOps) ZDiffWithScores(ctx context.Context, keys ...string) ([]ZItem, error) {
//	args := []any{"ZDIFF", len(keys)}
//	for _, key := range keys {
//		args = append(args, key)
//	}
//	args = append(args, "WITHSCORES")
//	return c.ZItemSlice(ctx, args...)
//}
//
//// ZIncrBy https://redis.io/commands/zincrby
//// Command: ZINCRBY key increment member
//// Bulk string reply: the new score of member
//// (a double precision floating point number), represented as string.
//func (c *orderedSetOps) ZIncrBy(ctx context.Context, key string, increment float64, member string) (float64, error) {
//	args := []any{"ZINCRBY", key, increment, member}
//	return toFloat64(c.driver.Exec((ctx, args...)
//}
//
//// ZInter https://redis.io/commands/zinter
//// Command: ZINTER numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX] [WITHSCORES]
//// Array reply: the result of intersection.
//func (c *orderedSetOps) ZInter(ctx context.Context, args ...any) ([]string, error) {
//	args = append([]any{"ZINTER"}, args...)
//	return toString(c.driver.Exec(Slice(ctx, args...)
//}
//
//// ZInterWithScores https://redis.io/commands/zinter
//// Command: ZINTER numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX] [WITHSCORES]
//// Array reply: the result of intersection.
//func (c *orderedSetOps) ZInterWithScores(ctx context.Context, args ...any) ([]ZItem, error) {
//	args = append([]any{"ZINTER"}, args...)
//	args = append(args, "WITHSCORES")
//	return c.ZItemSlice(ctx, args...)
//}
//
//// ZLexCount https://redis.io/commands/zlexcount
//// Command: ZLEXCOUNT key min max
//// Integer reply: the number of elements in the specified score range.
//func (c *orderedSetOps) ZLexCount(ctx context.Context, key, min, max string) (int64, error) {
//	args := []any{"ZLEXCOUNT", key, min, max}
//	return toInt64(c.driver.Exec(ctx, args...)
//}
//
//// ZMScore https://redis.io/commands/zmscore
//// Command: ZMSCORE key member [member ...]
//// Array reply: list of scores or nil associated with the specified member
//// values (a double precision floating point number), represented as strings.
//func (c *orderedSetOps) ZMScore(ctx context.Context, key string, members ...string) ([]float64, error) {
//	args := []any{"ZMSCORE", key}
//	for _, member := range members {
//		args = append(args, member)
//	}
//	return toFloat64(c.driver.Exec(Slice(ctx, args...)
//}
//
//// ZPopMax https://redis.io/commands/zpopmax
//// Command: ZPOPMAX key [count]
//// Array reply: list of popped elements and scores.
//func (c *orderedSetOps) ZPopMax(ctx context.Context, key string) ([]ZItem, error) {
//	args := []any{"ZPOPMAX", key}
//	return c.ZItemSlice(ctx, args...)
//}
//
//// ZPopMaxN https://redis.io/commands/zpopmax
//// Command: ZPOPMAX key [count]
//// Array reply: list of popped elements and scores.
//func (c *orderedSetOps) ZPopMaxN(ctx context.Context, key string, count int64) ([]ZItem, error) {
//	args := []any{"ZPOPMAX", key, count}
//	return c.ZItemSlice(ctx, args...)
//}
//
//// ZPopMin https://redis.io/commands/zpopmin
//// Command: ZPOPMIN key [count]
//// Array reply: list of popped elements and scores.
//func (c *orderedSetOps) ZPopMin(ctx context.Context, key string) ([]ZItem, error) {
//	args := []any{"ZPOPMIN", key}
//	return c.ZItemSlice(ctx, args...)
//}
//
//// ZPopMinN https://redis.io/commands/zpopmin
//// Command: ZPOPMIN key [count]
//// Array reply: list of popped elements and scores.
//func (c *orderedSetOps) ZPopMinN(ctx context.Context, key string, count int64) ([]ZItem, error) {
//	args := []any{"ZPOPMIN", key, count}
//	return c.ZItemSlice(ctx, args...)
//}
//
//// ZRandMember https://redis.io/commands/zrandmember
//// Command: ZRANDMEMBER key [count [WITHSCORES]]
//// Bulk Reply with the randomly selected element, or nil when key does not exist.
//func (c *orderedSetOps) ZRandMember(ctx context.Context, key string) (string, error) {
//	args := []any{"ZRANDMEMBER", key}
//	return toString(c.driver.Exec((ctx, args...)
//}
//
//// ZRandMemberN https://redis.io/commands/zrandmember
//// Command: ZRANDMEMBER key [count [WITHSCORES]]
//// Bulk Reply with the randomly selected element, or nil when key does not exist.
//func (c *orderedSetOps) ZRandMemberN(ctx context.Context, key string, count int) ([]string, error) {
//	args := []any{"ZRANDMEMBER", key, count}
//	return toString(c.driver.Exec(Slice(ctx, args...)
//}
//
//// ZRandMemberWithScores https://redis.io/commands/zrandmember
//// Command: ZRANDMEMBER key [count [WITHSCORES]]
//// Bulk Reply with the randomly selected element, or nil when key does not exist.
//func (c *orderedSetOps) ZRandMemberWithScores(ctx context.Context, key string, count int) ([]ZItem, error) {
//	args := []any{"ZRANDMEMBER", key, count, "WITHSCORES"}
//	return c.ZItemSlice(ctx, args...)
//}
//
//// ZRange https://redis.io/commands/zrange
//// Command: ZRANGE key min max [BYSCORE|BYLEX] [REV] [LIMIT offset count] [WITHSCORES]
//// Array reply: list of elements in the specified range.
//func (c *orderedSetOps) ZRange(ctx context.Context, key string, start, stop int64, args ...any) ([]string, error) {
//	args = append([]any{"ZRANGE", key, start, stop}, args...)
//	return toString(c.driver.Exec(Slice(ctx, args...)
//}
//
//// ZRangeWithScores https://redis.io/commands/zrange
//// Command: ZRANGE key min max [BYSCORE|BYLEX] [REV] [LIMIT offset count] [WITHSCORES]
//// Array reply: list of elements in the specified range.
//func (c *orderedSetOps) ZRangeWithScores(ctx context.Context, key string, start, stop int64, args ...any) ([]ZItem, error) {
//	args = append([]any{"ZRANGE", key, start, stop}, args...)
//	args = append(args, "WITHSCORES")
//	return c.ZItemSlice(ctx, args...)
//}
//
//// ZRangeByLex https://redis.io/commands/zrangebylex
//// Command: ZRANGEBYLEX key min max [LIMIT offset count]
//// Array reply: list of elements in the specified score range.
//func (c *orderedSetOps) ZRangeByLex(ctx context.Context, key string, min, max string, args ...any) ([]string, error) {
//	args = append([]any{"ZRANGEBYLEX", key, min, max}, args...)
//	return toString(c.driver.Exec(Slice(ctx, args...)
//}
//
//// ZRangeByScore https://redis.io/commands/zrangebyscore
//// Command: ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT offset count]
//// Array reply: list of elements in the specified score range.
//func (c *orderedSetOps) ZRangeByScore(ctx context.Context, key string, min, max string, args ...any) ([]string, error) {
//	args = append([]any{"ZRANGEBYSCORE", key, min, max}, args...)
//	return toString(c.driver.Exec(Slice(ctx, args...)
//}
//
//// ZRank https://redis.io/commands/zrank
//// Command: ZRANK key member
//// If member exists in the sorted set, Integer reply: the rank of member.
//// If member does not exist in the sorted set or key does not exist, Bulk string reply: nil.
//func (c *orderedSetOps) ZRank(ctx context.Context, key, member string) (int64, error) {
//	args := []any{"ZRANK", key, member}
//	return toInt64(c.driver.Exec(ctx, args...)
//}
//
//// ZRem https://redis.io/commands/zrem
//// Command: ZREM key member [member ...]
//// Integer reply, The number of members removed from the sorted set, not including non existing members.
//func (c *orderedSetOps) ZRem(ctx context.Context, key string, members ...any) (int64, error) {
//	args := []any{"ZREM", key}
//	for _, member := range members {
//		args = append(args, member)
//	}
//	return toInt64(c.driver.Exec(ctx, args...)
//}
//
//// ZRemRangeByLex https://redis.io/commands/zremrangebylex
//// Command: ZREMRANGEBYLEX key min max
//// Integer reply: the number of elements removed.
//func (c *orderedSetOps) ZRemRangeByLex(ctx context.Context, key, min, max string) (int64, error) {
//	args := []any{"ZREMRANGEBYLEX", key, min, max}
//	return toInt64(c.driver.Exec(ctx, args...)
//}
//
//// ZRemRangeByRank https://redis.io/commands/zremrangebyrank
//// Command: ZREMRANGEBYRANK key start stop
//// Integer reply: the number of elements removed.
//func (c *orderedSetOps) ZRemRangeByRank(ctx context.Context, key string, start, stop int64) (int64, error) {
//	args := []any{"ZREMRANGEBYRANK", key, start, stop}
//	return toInt64(c.driver.Exec(ctx, args...)
//}
//
//// ZRemRangeByScore https://redis.io/commands/zremrangebyscore
//// Command: ZREMRANGEBYSCORE key min max
//// Integer reply: the number of elements removed.
//func (c *orderedSetOps) ZRemRangeByScore(ctx context.Context, key, min, max string) (int64, error) {
//	args := []any{"ZREMRANGEBYSCORE", key, min, max}
//	return toInt64(c.driver.Exec(ctx, args...)
//}
//
//// ZRevRange https://redis.io/commands/zrevrange
//// Command: ZREVRANGE key start stop [WITHSCORES]
//// Array reply: list of elements in the specified range.
//func (c *orderedSetOps) ZRevRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
//	args := []any{"ZREVRANGE", key, start, stop}
//	return toString(c.driver.Exec(Slice(ctx, args...)
//}
//
//// ZRevRangeWithScores https://redis.io/commands/zrevrange
//// Command: ZREVRANGE key start stop [WITHSCORES]
//// Array reply: list of elements in the specified range.
//func (c *orderedSetOps) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) ([]string, error) {
//	args := []any{"ZREVRANGE", key, start, stop, "WITHSCORES"}
//	return toString(c.driver.Exec(Slice(ctx, args...)
//}
//
//// ZRevRangeByLex https://redis.io/commands/zrevrangebylex
//// Command: ZREVRANGEBYLEX key max min [LIMIT offset count]
//// Array reply: list of elements in the specified score range.
//func (c *orderedSetOps) ZRevRangeByLex(ctx context.Context, key string, min, max string, args ...any) ([]string, error) {
//	args = append([]any{"ZREVRANGEBYLEX", key, min, max}, args...)
//	return toString(c.driver.Exec(Slice(ctx, args...)
//}
//
//// ZRevRangeByScore https://redis.io/commands/zrevrangebyscore
//// Command: ZREVRANGEBYSCORE key max min [WITHSCORES] [LIMIT offset count]
//// Array reply: list of elements in the specified score range.
//func (c *orderedSetOps) ZRevRangeByScore(ctx context.Context, key string, min, max string, args ...any) ([]string, error) {
//	args = append([]any{"ZREVRANGEBYSCORE", key, min, max}, args...)
//	return toString(c.driver.Exec(Slice(ctx, args...)
//}
//
//// ZRevRank https://redis.io/commands/zrevrank
//// Command: ZREVRANK key member
//// If member exists in the sorted set, Integer reply: the rank of member.
//// If member does not exist in the sorted set or key does not exist, Bulk string reply: nil.
//func (c *orderedSetOps) ZRevRank(ctx context.Context, key, member string) (int64, error) {
//	args := []any{"ZREVRANK", key, member}
//	return toInt64(c.driver.Exec(ctx, args...)
//}
//
//// ZScore https://redis.io/commands/zscore
//// Command: ZSCORE key member
//// Bulk string reply: the score of member (a double precision floating point number), represented as string.
//func (c *orderedSetOps) ZScore(ctx context.Context, key, member string) (float64, error) {
//	args := []any{"ZSCORE", key, member}
//	return toFloat64(c.driver.Exec((ctx, args...)
//}
//
//// ZUnion https://redis.io/commands/zunion
//// Command: ZUNION numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX] [WITHSCORES]
//// Array reply: the result of union.
//func (c *orderedSetOps) ZUnion(ctx context.Context, args ...any) ([]string, error) {
//	args = append([]any{"ZUNION"}, args...)
//	return toString(c.driver.Exec(Slice(ctx, args...)
//}
//
//// ZUnionWithScores https://redis.io/commands/zunion
//// Command: ZUNION numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX] [WITHSCORES]
//// Array reply: the result of union.
//func (c *orderedSetOps) ZUnionWithScores(ctx context.Context, args ...any) ([]ZItem, error) {
//	args = append([]any{"ZUNION"}, args...)
//	args = append(args, "WITHSCORES")
//	return c.ZItemSlice(ctx, args...)
//}
//
//// ZUnionStore https://redis.io/commands/zunionstore
//// Command: ZUNIONSTORE destination numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX]
//// Integer reply: the number of elements in the resulting sorted set at destination.
//func (c *orderedSetOps) ZUnionStore(ctx context.Context, dest string, args ...any) (int64, error) {
//	args = append([]any{"ZUNIONSTORE", dest}, args...)
//	return toInt64(c.driver.Exec(ctx, args...)
//}
