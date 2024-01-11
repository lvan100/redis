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

func (c *myStringCmd) Set(ctx context.Context, key string, value any) *redis.StringCmdSet {
	return c.next.Set(ctx, key, value)
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
