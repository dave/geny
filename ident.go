package geny

import (
	"context"
	"io"
)

type ident struct {
	pkg    string
	name   string
	extras []Code
}

func Ident(name string) *ident {
	return &ident{
		name: name,
	}
}

func QualifiedIdent(pkg string, name string) *ident {
	return &ident{
		pkg:  pkg,
		name: name,
	}
}

func (i *ident) Build(ctx context.Context, w io.Writer) error {
	gc := FromContext(ctx)

	if gc.Package != i.pkg && i.pkg != "" {
		alias := gc.RegisterPackage(i.pkg)
		w.Write([]byte(alias))
		w.Write([]byte("."))
	}
	w.Write([]byte(i.name))

	for _, e := range i.extras {
		e.Build(ctx, w)
	}

	return nil

}

/*
func (i *ident) Call(code ...Code) *ident {
	i.extras = append(i.extras, call(code))
	return nil
}

type call []Code

func (c call) Build(ctx context.Context, w io.Writer) error {
	w.Write([]byte("("))
	commaSeparatedList(c).Build(ctx, w)
	w.Write([]byte(")"))
	return nil
}

func (i *ident) Index(code Code) *ident {
	i.extras = append(i.extras, &index{code: code})
	return nil
}

type index struct {
	code Code
}

func (i *index) Build(ctx context.Context, w io.Writer) error {
	w.Write([]byte("["))
	i.code.Build(ctx, w)
	w.Write([]byte("]"))
	return nil
}
*/
