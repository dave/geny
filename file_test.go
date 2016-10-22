package geny_test

import (
	"testing"

	. "github.com/davelondon/geny"
)

func TestFile(t *testing.T) {
	var code Code
	var expected string

	code = File()
	expected = `
package a`
	compare(t, "a", "a", code, expected)

}
