package evs

import (
	"testing"
)

func TestFrameString(t *testing.T) {
	frame := Frame{
		Line:     23,
		File:     "test.go",
		Function: "TestFrameString",
	}
	if "TestFrameString [test.go:23]" != frame.String() {
		t.Fatal("output was not equal")
	}
}
