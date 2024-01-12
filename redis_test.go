package redis_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/lvan100/redis"
)

type mockDriver struct{}

func (d *mockDriver) Exec(ctx context.Context, cmd string, args []any) (any, error) {
	fmt.Println("mockDriver", ctx, cmd, args)
	if cmd == "APPEND" {
		return int64(1), nil
	}
	return "OK", nil
}

type myStringOps struct {
	redis.StringOps
}

func (c *myStringOps) Set(ctx context.Context, key string, value any) *redis.CmdSet {
	fmt.Println("myStringOps", ctx, key, value)
	return c.StringOps.Set(ctx, key, value)
}

var client *redis.Client

func init() {
	var (
		err  error
		opts []redis.ClientOption
	)
	opts = append(opts, redis.CustomStringOps(func(ops redis.StringOps) redis.StringOps {
		return &myStringOps{StringOps: ops}
	}))
	client, err = redis.NewClient(&mockDriver{}, opts...)
	if err != nil {
		panic(err)
	}
}

func TestStringOps_Set(t *testing.T) {
	ctx := context.Background()
	{
		r, err := client.StringOps().Set(ctx, "a", "b").Ex(5).Xx().Get().Result()
		if err != nil {
			panic(err)
		}
		fmt.Println(r)
	}
	{
		r, err := client.StringOps().Append(ctx, "a", "b").Result()
		if err != nil {
			panic(err)
		}
		fmt.Println(r)
	}
}
