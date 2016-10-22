package geny

import (
	"context"
	"io"
)

type qualifiedIdent struct {
	pkg  string
	name string
}

func QualifiedIdent(pkg string, name string) *qualifiedIdent {
	return &qualifiedIdent{
		pkg:  pkg,
		name: name,
	}
}

func (i *qualifiedIdent) Build(ctx context.Context, w io.Writer) error {
	gc := FromContext(ctx)

	if gc.Package == i.pkg {
		w.Write([]byte(i.name))
		return nil
	}

	alias := gc.RegisterPackage(i.pkg)
	w.Write([]byte(alias))
	w.Write([]byte("."))
	w.Write([]byte(i.name))
	return nil

}
