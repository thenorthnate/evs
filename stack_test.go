//go:build whitebox

package evs

import (
	"strings"
	"testing"
)

func TestGetStack(t *testing.T) {
	stack := GetStack(0)
	stackStr := stack.String()
	if !strings.Contains(stackStr, "TestGetStack") {
		t.Fatal("stack string did not contain expected output")
	}
}
