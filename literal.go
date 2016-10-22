package geny

import (
	"context"
	"fmt"
	"io"
)

type literal struct {
	value interface{}
}

func Literal(value interface{}) *literal {
	return &literal{
		value: value,
	}
}

func (l *literal) Build(ctx context.Context, w io.Writer) error {
	fmt.Fprintf(w, "%#v", l.value)
	return nil
}
