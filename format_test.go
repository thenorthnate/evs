package evs

import (
	"errors"
	"testing"
)

func TestStandardFormatter(t *testing.T) {
	err := Error{
		Wraps: errors.New("bad error"),
		Stack: Stack{
			Frames: []Frame{{
				Line:     0,
				File:     "file.go",
				Function: "FunctionName",
			}},
		},
		Details: []Context{{
			Message: "oh no!",
			Location: Frame{
				Line:     1,
				File:     "file.go",
				Function: "SomeOtherFunctionName",
			},
		}},
		f: StandardFormatter{},
	}
	result := err.Error()
	expect := `bad error
[file.go:1] oh no!

With Stacktrace:
FunctionName [file.go:0]`
	if result != expect {
		t.Fatalf("Expected %v but got %v", expect, result)
	}
}
