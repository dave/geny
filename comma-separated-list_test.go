package geny

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/davelondon/ktest/assert"
)

func TestCommaSeparatedList_Build(t *testing.T) {

	var code Code
	var expected string

	code = commaSeparatedList([]Code{Ident("b")})
	expected = `b`
	compareRaw(t, "a", "a", code, expected)

	code = commaSeparatedList([]Code{Ident("b"), Ident("c")})
	expected = `b, c`
	compareRaw(t, "a", "a", code, expected)

}

func compareRaw(t *testing.T, path string, name string, code Code, expected string) {
	ctx := Context(context.Background(), path, name)
	buf := &bytes.Buffer{}
	code.Build(ctx, buf)
	source := buf.String()

	// trim any leading and trailing newlines
	expected = strings.Trim(expected, "\n")
	source = strings.Trim(source, "\n")

	if expected != source {
		fmt.Println(source)
	}
	assert.Equal(t, expected, source)
}
