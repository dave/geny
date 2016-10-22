package geny

import (
	"context"
	"io"
)

type identifier struct {
	name string
}

func Ident(name string) *identifier {
	return &identifier{
		name: name,
	}
}

func (i *identifier) Build(ctx context.Context, w io.Writer) error {
	w.Write([]byte(i.name))
	return nil

}
