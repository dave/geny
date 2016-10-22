package geny

import (
	"context"
	"io"
)

type param struct {
	identifier Code
	typ        Code
}

func Param(identifier Code, typ Code) *param {
	return &param{
		identifier: identifier,
		typ:        typ,
	}
}

func (p *param) Build(ctx context.Context, w io.Writer) error {
	p.identifier.Build(ctx, w)
	w.Write([]byte(" "))
	p.typ.Build(ctx, w)
	return nil

}
