package geny

import (
	"context"
	"io"
)

type fnc struct {
	ident  Code
	params []Code
	result []Code
	body   []Code
}

func FunctionDecl(ident Code) *fnc {
	return &fnc{
		ident: ident,
	}
}

func (f *fnc) Params(parameters ...Code) *fnc {
	f.params = parameters
	return f
}

func (f *fnc) Result(result ...Code) *fnc {
	f.result = result
	return f
}

func (f *fnc) Body(body ...Code) *fnc {
	f.body = body
	return f
}

func (f *fnc) Build(ctx context.Context, w io.Writer) error {

	ctx = Scope(ctx)

	// write the package
	w.Write([]byte("func "))
	f.ident.Build(ctx, w)
	w.Write([]byte("("))
	for i, param := range f.params {
		if i > 0 {
			w.Write([]byte(", "))
		}
		param.Build(ctx, w)
	}
	w.Write([]byte(") "))

	if len(f.result) > 0 {
		w.Write([]byte("("))
		for i, param := range f.result {
			if i > 0 {
				w.Write([]byte(", "))
			}
			param.Build(ctx, w)
		}
		w.Write([]byte(") "))
	}

	w.Write([]byte("{\n"))
	for _, c := range f.body {
		c.Build(ctx, w)
	}
	w.Write([]byte("}\n"))

	return nil

}
