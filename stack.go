package evs

import (
	"runtime"
)

const (
	startStackDepth = 10
	maxStackDepth   = 200
)

// Frame defines a single frame in a stack trace.
type Frame struct {
	Line     int
	File     string
	Function string
}

// CurrentFrame gets the location information for the code point where this function was called from (or
// anywhere up or down the stack from there depending on the skip value given.)
func CurrentFrame(skip int) Frame {
	skip++
	pc, file, line, _ := runtime.Caller(skip)
	function := runtime.FuncForPC(pc)
	return Frame{
		Line:     line,
		File:     file,
		Function: function.Name(),
	}
}

// Stack contains a stack trace made up of individual frames.
type Stack struct {
	Frames []Frame
}

// GetStack returns the full set of frames excluding the frames within the evs package
// assuming an appropriate value for Skip has been supplied. To get the stack excluding the
// call to [GetStack] itself (and everything beneath it), the value for skip should be 0.
func GetStack(skip int) Stack {
	skip++
	callerPCs := getCallerPCs(skip)
	callersFrames := runtime.CallersFrames(callerPCs)
	frames := []Frame{}
	for {
		frame, more := callersFrames.Next()
		frames = append(frames, Frame{
			Line:     frame.Line,
			File:     frame.File,
			Function: frame.Function,
		})
		// Check whether there are more frames to process after this one.
		if !more {
			break
		}
	}
	return Stack{Frames: frames}
}

func getCallerPCs(skip int) []uintptr {
	skip++
	skip++
	stackDepth := startStackDepth
	callerPCs := make([]uintptr, stackDepth)
	count := runtime.Callers(skip, callerPCs)
	if count == 0 {
		return []uintptr{}
	}

	for {
		if count < stackDepth || stackDepth == maxStackDepth {
			return callerPCs[:count]
		} else {
			stackDepth = stackDepth * 2
			if stackDepth > maxStackDepth {
				stackDepth = maxStackDepth
			}
			callerPCs = make([]uintptr, stackDepth)
			count = runtime.Callers(skip, callerPCs)
		}
	}
}
