package evs

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewError(t *testing.T) {
	err := newError(0)
	fmt.Println(err.Error())
	if !strings.Contains(err.Error(), "TestNewError") {
		t.Fatal("error did not contain expected output")
	}
}
