package evs

import (
	"fmt"
	"strings"
	"testing"
)

func TestFrameString(t *testing.T) {
	frame := Frame{
		Line:     23,
		File:     "test.go",
		Function: "TestFrameString",
	}
	if "TestFrameString [test.go:23]" != fmt.Sprintf("%v", frame) {
		t.Fatal("output was not equal")
	}
}

func TestGetStack(t *testing.T) {
	stack := GetStack(0)
	stackStr := fmt.Sprintf("%v", stack)
	fmt.Println(stackStr)
	if !strings.Contains(stackStr, "TestGetStack") {
		t.Fatal("stack string did not contain expected output")
	}
}

func doRecursion(count int) Stack {
	count--
	if count <= 0 {
		return GetStack(0)
	}
	return doRecursion(count)
}

func TestGetStackDeep(t *testing.T) {
	depth := startStackDepth + 2
	stack := doRecursion(depth)
	if len(stack.Frames) <= depth {
		t.Fatalf("expected stack to be at least %v in depth but was only %v", depth, len(stack.Frames))
	}
	stackStr := fmt.Sprintf("%v", stack)
	if !strings.Contains(stackStr, "TestGetStackDeep") {
		t.Fatal("stack string did not contain expected output")
	}
}

// func ExampleGetStack() {
// 	stack := GetStack(0)
// 	firstFrame := fmt.Sprintf("%s", stack.Frames[0])
// 	fmt.Println(strings.Contains(firstFrame, "[stack_test.go:50]"))
// 	// Output: true
// }
