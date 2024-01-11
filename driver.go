package redis

import (
	"context"
)

type Driver interface {
	Exec(ctx context.Context, cmd string, args []any) (any, error)
}
