package redis

import "context"

type SetOps interface {

	// SAdd https://redis.io/commands/sadd
	// Command: SADD key member [member ...]
	// Integer reply: the number of elements that were added to the set,
	// not including all the elements already present in the set.
	SAdd(ctx context.Context, key string, members ...any) (int64, error)

	// SCard https://redis.io/commands/scard
	// Command: SCARD key
	// Integer reply: the cardinality (number of elements) of the set,
	// or 0 if key does not exist.
	SCard(ctx context.Context, key string) (int64, error)

	// SDiff https://redis.io/commands/sdiff
	// Command: SDIFF key [key ...]
	// Array reply: list with members of the resulting set.
	SDiff(ctx context.Context, keys ...string) ([]string, error)

	// SDiffStore https://redis.io/commands/sdiffstore
	// Command: SDIFFSTORE destination key [key ...]
	// Integer reply: the number of elements in the resulting set.
	SDiffStore(ctx context.Context, destination string, keys ...string) (int64, error)

	// SInter https://redis.io/commands/sinter
	// Command: SINTER key [key ...]
	// Array reply: list with members of the resulting set.
	SInter(ctx context.Context, keys ...string) ([]string, error)

	// SInterStore https://redis.io/commands/sinterstore
	// Command: SINTERSTORE destination key [key ...]
	// Integer reply: the number of elements in the resulting set.
	SInterStore(ctx context.Context, destination string, keys ...string) (int64, error)

	// SIsMember https://redis.io/commands/sismember
	// Command: SISMEMBER key member
	// Integer reply: 1 if the element is a member of the set,
	// 0 if the element is not a member of the set, or if key does not exist.
	SIsMember(ctx context.Context, key string, member any) (int64, error)

	// SMembers https://redis.io/commands/smembers
	// Command: SMEMBERS key
	// Array reply: all elements of the set.
	SMembers(ctx context.Context, key string) ([]string, error)

	// SMIsMember https://redis.io/commands/smismember
	// Command: SMISMEMBER key member [member ...]
	// Array reply: list representing the membership of the given elements,
	// in the same order as they are requested.
	SMIsMember(ctx context.Context, key string, members ...any) ([]int64, error)

	// SMove https://redis.io/commands/smove
	// Command: SMOVE source destination member
	// Integer reply: 1 if the element is moved, 0 if the element
	// is not a member of source and no operation was performed.
	SMove(ctx context.Context, source, destination string, member any) (int64, error)

	// SPop https://redis.io/commands/spop
	// Command: SPOP key [count]
	// Bulk string reply: the removed member, or nil when key does not exist.
	SPop(ctx context.Context, key string) (string, error)

	// SPopN https://redis.io/commands/spop
	// Command: SPOP key [count]
	// Array reply: the removed members, or an empty array when key does not exist.
	SPopN(ctx context.Context, key string, count int64) ([]string, error)

	// SRandMember https://redis.io/commands/srandmember
	// Command: SRANDMEMBER key [count]
	// Returns a Bulk Reply with the randomly selected element,
	// or nil when key does not exist.
	SRandMember(ctx context.Context, key string) (string, error)

	// SRandMemberN https://redis.io/commands/srandmember
	// Command: SRANDMEMBER key [count]
	// Returns an array of elements, or an empty array when key does not exist.
	SRandMemberN(ctx context.Context, key string, count int64) ([]string, error)

	// SRem https://redis.io/commands/srem
	// Command: SREM key member [member ...]
	// Integer reply: the number of members that were removed from the set,
	// not including non-existing members.
	SRem(ctx context.Context, key string, members ...any) (int64, error)

	// SUnion https://redis.io/commands/sunion
	// Command: SUNION key [key ...]
	// Array reply: list with members of the resulting set.
	SUnion(ctx context.Context, keys ...string) ([]string, error)

	// SUnionStore https://redis.io/commands/sunionstore
	// Command: SUNIONSTORE destination key [key ...]
	// Integer reply: the number of elements in the resulting set.
	SUnionStore(ctx context.Context, destination string, keys ...string) (int64, error)
}

///////////////////////////////
//
//var _ SetOps = (*setOps)(nil)
//
//type setOps struct {
//	driver Driver
//}
//
//// SAdd https://redis.io/commands/sadd
//// Command: SADD key member [member ...]
//// Integer reply: the number of elements that were added to the set,
//// not including all the elements already present in the set.
//func (c *setOps) SAdd(ctx context.Context, key string, members ...any) (int64, error) {
//	args := []any{"SADD", key}
//	args = append(args, members...)
//	return toInt64(c.driver.Exec(ctx, args...)
//}
//
//// SCard https://redis.io/commands/scard
//// Command: SCARD key
//// Integer reply: the cardinality (number of elements) of the set,
//// or 0 if key does not exist.
//func (c *setOps) SCard(ctx context.Context, key string) (int64, error) {
//	args := []any{"SCARD", key}
//	return toInt64(c.driver.Exec(ctx, args...)
//}
//
//// SDiff https://redis.io/commands/sdiff
//// Command: SDIFF key [key ...]
//// Array reply: list with members of the resulting set.
//func (c *setOps) SDiff(ctx context.Context, keys ...string) ([]string, error) {
//	args := []any{"SDIFF"}
//	for _, key := range keys {
//		args = append(args, key)
//	}
//	return toString(c.driver.Exec(Slice(ctx, args...)
//}
//
//// SDiffStore https://redis.io/commands/sdiffstore
//// Command: SDIFFSTORE destination key [key ...]
//// Integer reply: the number of elements in the resulting set.
//func (c *setOps) SDiffStore(ctx context.Context, destination string, keys ...string) (int64, error) {
//	args := []any{"SDIFFSTORE", destination}
//	for _, key := range keys {
//		args = append(args, key)
//	}
//	return toInt64(c.driver.Exec(ctx, args...)
//}
//
//// SInter https://redis.io/commands/sinter
//// Command: SINTER key [key ...]
//// Array reply: list with members of the resulting set.
//func (c *setOps) SInter(ctx context.Context, keys ...string) ([]string, error) {
//	args := []any{"SINTER"}
//	for _, key := range keys {
//		args = append(args, key)
//	}
//	return toString(c.driver.Exec(Slice(ctx, args...)
//}
//
//// SInterStore https://redis.io/commands/sinterstore
//// Command: SINTERSTORE destination key [key ...]
//// Integer reply: the number of elements in the resulting set.
//func (c *setOps) SInterStore(ctx context.Context, destination string, keys ...string) (int64, error) {
//	args := []any{"SINTERSTORE", destination}
//	for _, key := range keys {
//		args = append(args, key)
//	}
//	return toInt64(c.driver.Exec(ctx, args...)
//}
//
//// SIsMember https://redis.io/commands/sismember
//// Command: SISMEMBER key member
//// Integer reply: 1 if the element is a member of the set,
//// 0 if the element is not a member of the set, or if key does not exist.
//func (c *setOps) SIsMember(ctx context.Context, key string, member any) (int64, error) {
//	args := []any{"SISMEMBER", key, member}
//	return toInt64(c.driver.Exec(ctx, args...)
//}
//
//// SMembers https://redis.io/commands/smembers
//// Command: SMEMBERS key
//// Array reply: all elements of the set.
//func (c *setOps) SMembers(ctx context.Context, key string) ([]string, error) {
//	args := []any{"SMEMBERS", key}
//	return toString(c.driver.Exec(Slice(ctx, args...)
//}
//
//// SMIsMember https://redis.io/commands/smismember
//// Command: SMISMEMBER key member [member ...]
//// Array reply: list representing the membership of the given elements,
//// in the same order as they are requested.
//func (c *setOps) SMIsMember(ctx context.Context, key string, members ...any) ([]int64, error) {
//	args := []any{"SMISMEMBER", key}
//	for _, member := range members {
//		args = append(args, member)
//	}
//	return c.IntSlice(ctx, args...)
//}
//
//// SMove https://redis.io/commands/smove
//// Command: SMOVE source destination member
//// Integer reply: 1 if the element is moved, 0 if the element
//// is not a member of source and no operation was performed.
//func (c *setOps) SMove(ctx context.Context, source, destination string, member any) (int64, error) {
//	args := []any{"SMOVE", source, destination, member}
//	return toInt64(c.driver.Exec(ctx, args...)
//}
//
//// SPop https://redis.io/commands/spop
//// Command: SPOP key [count]
//// Bulk string reply: the removed member, or nil when key does not exist.
//func (c *setOps) SPop(ctx context.Context, key string) (string, error) {
//	args := []any{"SPOP", key}
//	return toString(c.driver.Exec((ctx, args...)
//}
//
//// SPopN https://redis.io/commands/spop
//// Command: SPOP key [count]
//// Array reply: the removed members, or an empty array when key does not exist.
//func (c *setOps) SPopN(ctx context.Context, key string, count int64) ([]string, error) {
//	args := []any{"SPOP", key, count}
//	return toString(c.driver.Exec(Slice(ctx, args...)
//}
//
//// SRandMember https://redis.io/commands/srandmember
//// Command: SRANDMEMBER key [count]
//// Returns a Bulk Reply with the randomly selected element,
//// or nil when key does not exist.
//func (c *setOps) SRandMember(ctx context.Context, key string) (string, error) {
//	args := []any{"SRANDMEMBER", key}
//	return toString(c.driver.Exec((ctx, args...)
//}
//
//// SRandMemberN https://redis.io/commands/srandmember
//// Command: SRANDMEMBER key [count]
//// Returns an array of elements, or an empty array when key does not exist.
//func (c *setOps) SRandMemberN(ctx context.Context, key string, count int64) ([]string, error) {
//	args := []any{"SRANDMEMBER", key, count}
//	return toString(c.driver.Exec(Slice(ctx, args...)
//}
//
//// SRem https://redis.io/commands/srem
//// Command: SREM key member [member ...]
//// Integer reply: the number of members that were removed from the set,
//// not including non-existing members.
//func (c *setOps) SRem(ctx context.Context, key string, members ...any) (int64, error) {
//	args := []any{"SREM", key}
//	for _, member := range members {
//		args = append(args, member)
//	}
//	return toInt64(c.driver.Exec(ctx, args...)
//}
//
//// SUnion https://redis.io/commands/sunion
//// Command: SUNION key [key ...]
//// Array reply: list with members of the resulting set.
//func (c *setOps) SUnion(ctx context.Context, keys ...string) ([]string, error) {
//	args := []any{"SUNION"}
//	for _, key := range keys {
//		args = append(args, key)
//	}
//	return toString(c.driver.Exec(Slice(ctx, args...)
//}
//
//// SUnionStore https://redis.io/commands/sunionstore
//// Command: SUNIONSTORE destination key [key ...]
//// Integer reply: the number of elements in the resulting set.
//func (c *setOps) SUnionStore(ctx context.Context, destination string, keys ...string) (int64, error) {
//	args := []any{"SUNIONSTORE", destination}
//	for _, key := range keys {
//		args = append(args, key)
//	}
//	return toInt64(c.driver.Exec(ctx, args...)
//}
