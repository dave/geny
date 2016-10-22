package geny_test

import (
	"testing"

	. "kego.io/process/geny"
)

func TestFunctionDecl(t *testing.T) {
	var code Code
	var expected string

	// minimum funcion decl
	code = File(
		FunctionDecl(
			Ident("a"),
		),
	)
	expected = `
package a

func a() {
}`
	compare(t, "a", "a", code, expected)

	// params, result and body
	code = File(
		FunctionDecl(
			Ident("a"),
		).Params(
			Param(Ident("b"), Ident("c")),
			Param(Ident("d"), Ident("e")),
		).Result(
			Ident("f"),
		).Body(
			Ident("g"),
		),
	)
	expected = `
package a

func a(b c, d e) f {
	g
}`
	compare(t, "a", "a", code, expected)

	// multiple result parameters
	code = File(
		FunctionDecl(
			Ident("a"),
		).Result(
			Param(Ident("b"), Ident("c")),
			Param(Ident("d"), Ident("e")),
		),
	)
	expected = `
package a

func a() (b c, d e) {
}`
	compare(t, "a", "a", code, expected)
}
