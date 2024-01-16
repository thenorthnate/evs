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
		Details: []Detail{{
			Message: "oh no!",
			Location: Frame{
				Line:     1,
				File:     "file.go",
				Function: "SomeOtherFunctionName",
			},
		}},
		f: standardFormatter{},
	}
}

func TestStandardFormatter(t *testing.T) {
	err := getTestError()
	result := err.Error()
	expect := `*evs.Error: bad error
SomeOtherFunctionName [file.go:1] oh no!

With Stacktrace:
FunctionName [file.go:0]`
	if result != expect {
		t.Fatalf("Expected\n%v\nbut got\n%v", expect, result)
	}
}

func TestStandardFormatterNoStack(t *testing.T) {
	err := getTestError()
	err.Stack = Stack{}
	result := err.Error()
	expect := `*evs.Error: bad error
SomeOtherFunctionName [file.go:1] oh no!`
	if result != expect {
		t.Fatalf("Expected\n%v\nbut got\n%v", expect, result)
	}
}

func TestStandardFormatterNoStackOrWrapped(t *testing.T) {
	err := getTestError()
	err.Wraps = nil
	err.Stack = Stack{}
	result := err.Error()
	expect := `*evs.Error: oh no!`
	if result != expect {
		t.Fatalf("Expected\n%v\nbut got\n%v", expect, result)
	}
}
