package geny_test

import (
	"context"
	"testing"

	"bytes"

	"fmt"
	"go/format"
	"strings"

	"github.com/davelondon/ktest/require"
	"github.com/stretchr/testify/assert"
	. "kego.io/process/geny"
)

func TestContext(t *testing.T) {
	var code Code
	var expected string

	ctx := Context(context.Background(), "a.b/c", "c")
	gc := FromContext(ctx)
	gc.RegisterAnonymousPackage("d.e/f")
	code = File()
	expected = `
package c

import _ "d.e/f"`
	compareCtx(t, ctx, code, expected)

	ctx = Context(context.Background(), "a.b/c", "c")
	gc = FromContext(ctx)
	gc.RegisterAnonymousPackage("d.e/f")
	code = File(
		VarDecl(Ident("a")).Type(QualifiedIdent("d.e/f", "g")),
		VarDecl(Ident("b")).Type(QualifiedIdent("d.e/f", "h")),
	)
	expected = `
package c

import f "d.e/f"

var a f.g
var b f.h`
	compareCtx(t, ctx, code, expected)

}

func compareCtx(t *testing.T, ctx context.Context, code Code, expected string) {
	buf := &bytes.Buffer{}
	code.Build(ctx, buf)

	// format the source
	sourceBytes, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Println(buf.String())
	}
	require.NoError(t, err)
	source := string(sourceBytes)

	// trim any leading and trailing newlines
	expected = strings.Trim(expected, "\n")
	source = strings.Trim(source, "\n")

	if expected != source {
		fmt.Println(source)
	}
	assert.Equal(t, expected, source)
}
