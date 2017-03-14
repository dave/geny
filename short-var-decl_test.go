package geny_test

import (
	"testing"

	. "github.com/dave/geny"
)

func TestShortVarDecl(t *testing.T) {
	var code Code
	var expected string

	// uses last part of url
	code = File(
		FunctionDecl(Ident("b")).Body(
			ShortVarDecl(Ident("c")).Equals(Ident("d")),
		),
	)
	expected = `
package a

func b() {
	c := d
}`
	compare(t, "a", "a", code, expected)

}
