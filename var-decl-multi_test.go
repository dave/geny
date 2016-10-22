package geny_test

import (
	"testing"

	. "github.com/davelondon/geny"
)

func TestVarDeclMulti(t *testing.T) {
	var code Code
	var expected string

	code = File(
		VarDeclMulti(
			VarDecl(Ident("a")).Type(Ident("b")),
		),
	)
	expected = `
package a

var a b`
	compare(t, "a", "a", code, expected)

	code = File(
		VarDeclMulti(
			VarDecl(Ident("a")).Type(Ident("b")),
			VarDecl(Ident("c")).Type(Ident("d")),
		),
	)
	expected = `
package a

var (
	a b
	c d
)`
	compare(t, "a", "a", code, expected)

	code = File(
		VarDeclMulti(
			VarDecl(Ident("a")).Type(Ident("b")),
			VarDecl(Ident("c"), Ident("d")).Type(Ident("e")).Equals(Ident("f")),
		),
	)
	expected = `
package a

var (
	a    b
	c, d e = f
)`
	compare(t, "a", "a", code, expected)
}
