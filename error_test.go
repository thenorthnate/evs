package evs

import (
	"strings"
	"testing"
)

func TestNewError(t *testing.T) {
	err := newError[Std](0)
	if !strings.Contains(err.Error(), "TestNewError") {
		t.Fatal("error did not contain expected output")
	}
}
