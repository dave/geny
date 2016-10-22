package geny_test

import (
	"testing"

	. "kego.io/process/geny"
)

func TestParam(t *testing.T) {
	var code Code
	var expected string

	code = File(
		FunctionDecl(Ident("b")).Params(Param(Ident("c"), Ident("d"))),
	)
	expected = `
package a

func b(c d) {
}`
	compare(t, "a", "a", code, expected)

}
