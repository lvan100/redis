package redis_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lvan100/redis"
)

type mockDriver struct{}

func (d *mockDriver) Exec(ctx context.Context, cmd string, args []any) (any, error) {
	fmt.Println(ctx, cmd, args)
	return "OK", nil
}

type myStringCmd struct {
	next redis.StringCmd
}

func (c *myStringCmd) Append(ctx context.Context, key, value string) (int64, error) {
	return c.next.Append(ctx, key, value)
}

func (c *myStringCmd) Decr(ctx context.Context, key string) (int64, error) {
	return c.next.Decr(ctx, key)
}

func (c *myStringCmd) DecrBy(ctx context.Context, key string, decrement int64) (int64, error) {
	return c.next.DecrBy(ctx, key, decrement)
}

func (c *myStringCmd) Get(ctx context.Context, key string) (string, error) {
	return c.next.Get(ctx, key)
}

func (c *myStringCmd) GetDel(ctx context.Context, key string) (string, error) {
	return c.next.GetDel(ctx, key)
}

func (c *myStringCmd) GetEx(ctx context.Context, key string, args ...any) *redis.StringCmdGet {
	return c.next.GetEx(ctx, key, args...)
}

func (c *myStringCmd) GetRange(ctx context.Context, key string, start, end int64) (string, error) {
	return c.next.GetRange(ctx, key, start, end)
}

func (c *myStringCmd) GetSet(ctx context.Context, key string, value any) (string, error) {
	return c.next.GetSet(ctx, key, value)
}

func (c *myStringCmd) Incr(ctx context.Context, key string) (int64, error) {
	return c.next.Incr(ctx, key)
}

func (c *myStringCmd) IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	return c.next.IncrBy(ctx, key, value)
}

func (c *myStringCmd) IncrByFloat(ctx context.Context, key string, value float64) (float64, error) {
	return c.next.IncrByFloat(ctx, key, value)
}

func (c *myStringCmd) MGet(ctx context.Context, keys ...string) ([]any, error) {
	return c.next.MGet(ctx, keys...)
}

func (c *myStringCmd) MSet(ctx context.Context, args ...any) (string, error) {
	return c.next.MSet(ctx, args...)
}

func (c *myStringCmd) MSetNX(ctx context.Context, args ...any) (int64, error) {
	return c.next.MSetNX(ctx, args...)
}

func (c *myStringCmd) PSetEX(ctx context.Context, key string, value any, expire int64) (string, error) {
	return c.next.PSetEX(ctx, key, value, expire)
}

func (c *myStringCmd) Set(ctx context.Context, key string, value any) *redis.StringCmdSet {
	return c.next.Set(ctx, key, value)
}

func (c *myStringCmd) SetEX(ctx context.Context, key string, value any, expire int64) (string, error) {
	return c.next.SetEX(ctx, key, value, expire)
}

func (c *myStringCmd) SetNX(ctx context.Context, key string, value any) (int64, error) {
	return c.next.SetNX(ctx, key, value)
}

func (c *myStringCmd) SetRange(ctx context.Context, key string, offset int64, value string) (int64, error) {
	return c.next.SetRange(ctx, key, offset, value)
}

func (c *myStringCmd) StrLen(ctx context.Context, key string) (int64, error) {
	return c.next.StrLen(ctx, key)
}

var client *redis.Client

func init() {
	var (
		err  error
		opts []redis.ClientOption
	)
	opts = append(opts, redis.CustomStringCmd(func(cmd redis.StringCmd) redis.StringCmd {
		return &myStringCmd{next: cmd}
	}))
	client, err = redis.NewClient(&mockDriver{}, opts...)
	if err != nil {
		panic(err)
	}
}

func TestStringCmd_Set(t *testing.T) {
	ctx := context.Background()
	r, err := client.StringCmd().Set(ctx, "a", "b").Ex(5).Send()
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}
