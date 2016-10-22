package geny

import (
	"context"
	"io"
)

type ident struct {
	pkg  string
	name string
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

	if gc.Package == i.pkg || i.pkg == "" {
		w.Write([]byte(i.name))
		return nil
	}

	alias := gc.RegisterPackage(i.pkg)
	w.Write([]byte(alias))
	w.Write([]byte("."))
	w.Write([]byte(i.name))
	return nil

}
