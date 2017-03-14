package geny_test

import (
	"testing"

	. "github.com/dave/geny"
)

func TestPredeclaredIdent_Build(t *testing.T) {
	var code Code
	var expected string

	code = File(
		VarDecl(Ident("a")).Type(Bool()),
		VarDecl(Ident("b")).Type(Byte()),
		VarDecl(Ident("c")).Type(Complex64()),
		VarDecl(Ident("d")).Type(Complex128()),
		VarDecl(Ident("e")).Type(Error()),
		VarDecl(Ident("f")).Type(Float32()),
		VarDecl(Ident("g")).Type(Float64()),
		VarDecl(Ident("h")).Type(Int()),
		VarDecl(Ident("i")).Type(Int8()),
		VarDecl(Ident("j")).Type(Int16()),
		VarDecl(Ident("k")).Type(Int32()),
		VarDecl(Ident("l")).Type(Int64()),
		VarDecl(Ident("m")).Type(Rune()),
		VarDecl(Ident("n")).Type(String()),
		VarDecl(Ident("o")).Type(Uint()),
		VarDecl(Ident("p")).Type(Uint8()),
		VarDecl(Ident("q")).Type(Uint16()),
		VarDecl(Ident("r")).Type(Uint32()),
		VarDecl(Ident("s")).Type(Uint64()),
		VarDecl(Ident("t")).Type(Uintptr()),
		VarDecl(Ident("u")).Type(True()),
		VarDecl(Ident("v")).Type(False()),
		VarDecl(Ident("w")).Type(Iota()),
		VarDecl(Ident("x")).Type(Nil()),
	)
	expected = `
package a

var a bool
var b byte
var c complex64
var d complex128
var e error
var f float32
var g float64
var h int
var i int8
var j int16
var k int32
var l int64
var m rune
var n string
var o uint
var p uint8
var q uint16
var r uint32
var s uint64
var t uintptr
var u true
var v false
var w iota
var x nil`
	compare(t, "a", "a", code, expected)

}
