package geny

import (
	"context"
	"io"
)

type shortVarDecl struct {
	identifiers Code
	expressions Code
}

func ShortVarDecl(identifiers ...Code) *shortVarDecl {
	return &shortVarDecl{
		identifiers: commaSeparatedList(identifiers),
	}
}

func (s *shortVarDecl) Equals(expressions ...Code) *shortVarDecl {
	s.expressions = commaSeparatedList(expressions)
	return s
}

func (s *shortVarDecl) Build(ctx context.Context, w io.Writer) error {
	if s.expressions == nil {
		panic("ShortVarDecl must specify expression with Equals method")
	}
	s.identifiers.Build(ctx, w)
	w.Write([]byte(" := "))
	s.expressions.Build(ctx, w)
	w.Write([]byte("\n"))
	return nil
}
