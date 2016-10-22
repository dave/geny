package geny

import (
	"context"
	"io"
)

type varDecl struct {
	identifiers Code
	typ         Code
	expressions Code
	multi       bool
}

func VarDecl(identifiers ...Code) *varDecl {
	return &varDecl{
		identifiers: commaSeparatedList(identifiers),
	}
}

func (s *varDecl) Type(typ Code) *varDecl {
	s.typ = typ
	return s
}

func (s *varDecl) Equals(expressions ...Code) *varDecl {
	s.expressions = commaSeparatedList(expressions)
	return s
}

func (s *varDecl) Build(ctx context.Context, w io.Writer) error {
	if !s.multi {
		w.Write([]byte("var "))
	}
	s.identifiers.Build(ctx, w)
	if s.typ != nil {
		w.Write([]byte(" "))
		s.typ.Build(ctx, w)
	}
	if s.expressions != nil {
		w.Write([]byte(" = "))
		s.expressions.Build(ctx, w)
	}
	w.Write([]byte("\n"))
	return nil
}
