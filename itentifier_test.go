package geny_test

import (
	"testing"

	. "kego.io/process/geny"
)

func TestIdent(t *testing.T) {
	var code Code
	var expected string

	code = File(
		VarDecl(Ident("b")).Type(Ident("c")),
	)
	expected = `
package a

var b c`
	compare(t, "a", "a", code, expected)

}
