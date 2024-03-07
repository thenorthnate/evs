package evs

import (
	"errors"
	"testing"
)

func getTestError() Error {
	return Error{
		Wraps: errors.New("bad error"),
		Stack: Stack{
			Frames: []Frame{{
				Line:     0,
				File:     "file.go",
				Function: "FunctionName",
			}},
		},
		Details: []string{"oh no!"},
		f:       textFormatter{},
	}
}

func TestTextFormatter(t *testing.T) {
	err := getTestError()
	result := err.Error()
	expect := `bad error
[oh no!]

With Stacktrace:
FunctionName [file.go:0]`
	if result != expect {
		t.Fatalf("Expected\n%v\nbut got\n%v", expect, result)
	}
}

func TestTextFormatterNoStack(t *testing.T) {
	err := getTestError()
	err.Stack = Stack{}
	result := err.Error()
	expect := `bad error
[oh no!]`
	if result != expect {
		t.Fatalf("Expected\n%v\nbut got\n%v", expect, result)
	}
}

func TestTextFormatterNoStackOrWrapped(t *testing.T) {
	err := getTestError()
	err.Wraps = nil
	err.Stack = Stack{}
	result := err.Error()
	expect := `[oh no!]`
	if result != expect {
		t.Fatalf("Expected\n%v\nbut got\n%v", expect, result)
	}
}
