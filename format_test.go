package evs

import (
	"encoding/json"
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
		f: textFormatter{},
	}
}

func TestTextFormatter(t *testing.T) {
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

func TestTextFormatterNoStack(t *testing.T) {
	err := getTestError()
	err.Stack = Stack{}
	result := err.Error()
	expect := `*evs.Error: bad error
SomeOtherFunctionName [file.go:1] oh no!`
	if result != expect {
		t.Fatalf("Expected\n%v\nbut got\n%v", expect, result)
	}
}

func TestTextFormatterNoStackOrWrapped(t *testing.T) {
	err := getTestError()
	err.Wraps = nil
	err.Stack = Stack{}
	result := err.Error()
	expect := `*evs.Error: oh no!`
	if result != expect {
		t.Fatalf("Expected\n%v\nbut got\n%v", expect, result)
	}
}

func TestJSONFormatter(t *testing.T) {
	err := New("something sad happened").
		Fmt(JSONFormatter()).
		Err()
	result := struct {
		Wraps   string
		Stack   Stack
		Details []Detail
	}{}
	if err := json.Unmarshal([]byte(err.Error()), &result); err != nil {
		t.Fatalf("encountered unexpected error: %v", err)
	}
	if len(result.Details) != 1 {
		t.Fatalf("expected a single set of details but got %v", len(result.Details))
	}
}

func TestJSONFormatter_From(t *testing.T) {
	err := From(errors.New("something sad happened")).
		Fmt(JSONFormatter()).
		Err()
	result := struct {
		Wraps   string
		Stack   Stack
		Details []Detail
	}{}
	if err := json.Unmarshal([]byte(err.Error()), &result); err != nil {
		t.Fatalf("encountered unexpected error: %v", err)
	}
	if len(result.Details) != 0 {
		t.Fatalf("expected a single set of details but got %v", len(result.Details))
	}
}
