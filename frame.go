package evs

import (
	"fmt"
	"strings"
)

// Frame defines a single frame in a stack trace.
type Frame struct {
	Line     int
	File     string
	Function string
}

// String implements the [fmt.Stringer] interface.
func (frame Frame) String() string {
	fileParts := strings.Split(frame.File, "/")
	return fmt.Sprintf("%v [%v:%v]", frame.Function, fileParts[len(fileParts)-1], frame.Line)
}
