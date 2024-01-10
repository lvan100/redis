package redis

import (
	"bytes"
	"context"
)

type Driver interface {
	Exec(ctx context.Context, buf *bytes.Buffer) (any, error)
}
