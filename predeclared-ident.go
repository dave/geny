package geny

import (
	"context"
	"io"
)

type predeclaredIdent struct {
	ident string
}

func Bool() *predeclaredIdent       { return &predeclaredIdent{ident: "bool"} }
func Byte() *predeclaredIdent       { return &predeclaredIdent{ident: "byte"} }
func Complex64() *predeclaredIdent  { return &predeclaredIdent{ident: "complex64"} }
func Complex128() *predeclaredIdent { return &predeclaredIdent{ident: "complex128"} }
func Error() *predeclaredIdent      { return &predeclaredIdent{ident: "error"} }
func Float32() *predeclaredIdent    { return &predeclaredIdent{ident: "float32"} }
func Float64() *predeclaredIdent    { return &predeclaredIdent{ident: "float64"} }
func Int() *predeclaredIdent        { return &predeclaredIdent{ident: "int"} }
func Int8() *predeclaredIdent       { return &predeclaredIdent{ident: "int8"} }
func Int16() *predeclaredIdent      { return &predeclaredIdent{ident: "int16"} }
func Int32() *predeclaredIdent      { return &predeclaredIdent{ident: "int32"} }
func Int64() *predeclaredIdent      { return &predeclaredIdent{ident: "int64"} }
func Rune() *predeclaredIdent       { return &predeclaredIdent{ident: "rune"} }
func String() *predeclaredIdent     { return &predeclaredIdent{ident: "string"} }
func Uint() *predeclaredIdent       { return &predeclaredIdent{ident: "uint"} }
func Uint8() *predeclaredIdent      { return &predeclaredIdent{ident: "uint8"} }
func Uint16() *predeclaredIdent     { return &predeclaredIdent{ident: "uint16"} }
func Uint32() *predeclaredIdent     { return &predeclaredIdent{ident: "uint32"} }
func Uint64() *predeclaredIdent     { return &predeclaredIdent{ident: "uint64"} }
func Uintptr() *predeclaredIdent    { return &predeclaredIdent{ident: "uintptr"} }
func True() *predeclaredIdent       { return &predeclaredIdent{ident: "true"} }
func False() *predeclaredIdent      { return &predeclaredIdent{ident: "false"} }
func Iota() *predeclaredIdent       { return &predeclaredIdent{ident: "iota"} }
func Nil() *predeclaredIdent        { return &predeclaredIdent{ident: "nil"} }

func (i *predeclaredIdent) Build(ctx context.Context, w io.Writer) error {
	w.Write([]byte(i.ident))
	return nil
}
