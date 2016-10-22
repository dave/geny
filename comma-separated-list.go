package geny

import (
	"context"
	"io"
)

type commaSeparatedList []Code

func (l commaSeparatedList) Build(ctx context.Context, w io.Writer) error {
	for i, c := range l {
		if i > 0 {
			w.Write([]byte(", "))
		}
		c.Build(ctx, w)
	}
	return nil

}
