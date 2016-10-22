package geny_test

import (
	"context"
	"testing"

	"bytes"

	"fmt"

	"go/format"

	"strings"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"
	. "kego.io/process/geny"
)

func TestGeny(t *testing.T) {
	ctx := Context(context.Background(), "a.b/c", "c")
	f := File(
		FunctionDecl(Ident("Foo")).Params(
			Param(Ident("f"), QualifiedIdent("a.b/c", "Foo")),
		).Result(
			Error(),
		).Body(
			ShortVarDecl(Ident("a")).Equals(QualifiedIdent("a.b/d", "Bar")),
			ShortVarDecl(
				Ident("a"),
				Ident("c"),
			).Equals(
				QualifiedIdent("a.b/c", "Bar"),
				QualifiedIdent("a.b/e", "Baz"),
			),
			VarDecl(Ident("b")).Equals(Ident("c")),
			VarDecl(Ident("c")).Type(Uint32()),
			VarDeclMulti(
				VarDecl(Ident("e")).Equals(Ident("f")),
				VarDecl(Ident("c")).Type(QualifiedIdent("context", "Context")).Equals(Nil()),
			),
			ShortVarDecl(Ident("a")).Equals(Literal("foo")),
		),
	)
	buf := &bytes.Buffer{}
	f.Build(ctx, buf)
	source, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Println(buf.String())
	}
	require.NoError(t, err)
	fmt.Println(string(source))
}

func compare(t *testing.T, path string, name string, code Code, expected string) {
	ctx := Context(context.Background(), path, name)
	buf := &bytes.Buffer{}
	code.Build(ctx, buf)

	// format the source
	sourceBytes, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Println(buf.String())
	}
	require.NoError(t, err)
	source := string(sourceBytes)

	// trim any leading and trailing newlines
	expected = strings.Trim(expected, "\n")
	source = strings.Trim(source, "\n")

	if expected != source {
		fmt.Println(source)
	}
	assert.Equal(t, expected, source)
}
