package geny_test

import (
	"testing"

	. "kego.io/process/geny"
)

func TestGuessPackageNames(t *testing.T) {
	var code Code
	var expected string

	// uses last part of url
	code = File(
		VarDecl(Ident("_")).Type(QualifiedIdent("b.c/d/e", "z")),
	)
	expected = `
package a

import e "b.c/d/e"

var _ e.z`
	compare(t, "a", "a", code, expected)

	// tolerates trailing slash and removes
	code = File(
		VarDecl(Ident("_")).Type(QualifiedIdent("b.c/d/e/", "z")),
	)
	expected = `
package a

import e "b.c/d/e"

var _ e.z`
	compare(t, "a", "a", code, expected)

	// uses last hyphen separated part
	code = File(
		VarDecl(Ident("_")).Type(QualifiedIdent("b.c/d/e-f-g", "z")),
	)
	expected = `
package a

import g "b.c/d/e-f-g"

var _ g.z`
	compare(t, "a", "a", code, expected)

	// uses first dot separated part
	code = File(
		VarDecl(Ident("_")).Type(QualifiedIdent("b.c/d/e.f.g", "z")),
	)
	expected = `
package a

import e "b.c/d/e.f.g"

var _ e.z`
	compare(t, "a", "a", code, expected)

	// hyphen bit is done first
	code = File(
		VarDecl(Ident("_")).Type(QualifiedIdent("b.c/d/e-f.g-h", "z")),
	)
	expected = `
package a

import h "b.c/d/e-f.g-h"

var _ h.z`
	compare(t, "a", "a", code, expected)

	// no imports
	code = File()
	expected = `
package c`
	compare(t, "a.b/c", "c", code, expected)

	// local import
	code = File(
		VarDecl(Ident("_")).Type(QualifiedIdent("a.b/c", "d")),
	)
	expected = `
package c

var _ d`
	compare(t, "a.b/c", "c", code, expected)

	// external import
	code = File(
		VarDecl(Ident("_")).Type(QualifiedIdent("e.f/g", "h")),
	)
	expected = `
package c

import g "e.f/g"

var _ g.h`
	compare(t, "a.b/c", "c", code, expected)

	// auto rename conflicting packages
	code = File(
		VarDecl(Ident("_")).Type(QualifiedIdent("e.f/g", "h")),
		VarDecl(Ident("_")).Type(QualifiedIdent("i.j/g", "k")),
	)
	expected = `
package c

import (
	g "e.f/g"
	g1 "i.j/g"
)

var _ g.h
var _ g1.k`
	compare(t, "a.b/c", "c", code, expected)
}
