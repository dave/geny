package geny

import (
	"context"
	"io"
)

type varDeclMulti struct {
	decls []*varDecl
}

func VarDeclMulti(decls ...*varDecl) *varDeclMulti {
	return &varDeclMulti{
		decls: decls,
	}
}

func (s *varDeclMulti) Build(ctx context.Context, w io.Writer) error {

	if len(s.decls) == 1 {
		s.decls[0].Build(ctx, w)
		return nil
	}

	w.Write([]byte("var ("))
	for _, d := range s.decls {
		d.multi = true
		d.Build(ctx, w)
	}
	w.Write([]byte(")\n"))
	return nil
}
