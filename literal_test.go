package geny_test

import (
	"testing"

	. "kego.io/process/geny"
)

func TestLiteral(t *testing.T) {
	var code Code
	var expected string

	code = File(
		VarDecl(Ident("a")).Equals(Literal("b")),
	)
	expected = `
package a

var a = "b"`
	compare(t, "a", "a", code, expected)

	code = File(
		VarDecl(Ident("a")).Equals(Literal(1.1)),
	)
	expected = `
package a

var a = 1.1`
	compare(t, "a", "a", code, expected)

	code = File(
		VarDecl(Ident("a")).Equals(Literal(true)),
	)
	expected = `
package a

var a = true`
	compare(t, "a", "a", code, expected)

}
