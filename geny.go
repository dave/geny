package geny

import (
	"context"
	"io"
)

type Code interface {
	Build(ctx context.Context, w io.Writer) error
}
