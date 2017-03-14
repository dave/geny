package geny_test

import (
	"testing"

	. "github.com/dave/geny"
)

func TestVarDecl(t *testing.T) {
	var code Code
	var expected string

	code = File(
		VarDecl(Ident("a")).Type(Ident("b")),
	)
	expected = `
package a

var a b`
	compare(t, "a", "a", code, expected)

	code = File(
		VarDecl(Ident("a")).Equals(Ident("b")),
	)
	expected = `
package a

var a = b`
	compare(t, "a", "a", code, expected)

	code = File(
		VarDecl(Ident("a")).Type(Ident("b")).Equals(Ident("c")),
	)
	expected = `
package a

var a b = c`
	compare(t, "a", "a", code, expected)

	code = File(
		VarDecl(Ident("a"), Ident("b")).Equals(Ident("c")),
	)
	expected = `
package a

var a, b = c`
	compare(t, "a", "a", code, expected)

	code = File(
		VarDecl(Ident("a"), Ident("b")).Equals(Ident("c"), Ident("d")),
	)
	expected = `
package a

var a, b = c, d`
	compare(t, "a", "a", code, expected)

}
