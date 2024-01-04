package evs

import (
	"errors"
	"fmt"
)

const (
	initialSkip = 1
)

var (
	// IncludeStack is used to determine whether or not a stacktrace should be captured with
	// new errors. By default it is set to true.
	IncludeStack = true
	// SurfaceLevel controls how [From] operates. By default, if the error provided to [From] is an [Error]
	// the given error is not "wrapped" but instead combined with any additional details you give it. If the
	// error is something else, but contains an [Error] that can be retreived via "Unwrap" it will
	// remain invisible. By setting [SurfaceLevel] to false, that means that using [From] will inspect the
	// full error stack via a call to [errors.As] and will drop other higher order errors if they exist.
	SurfaceLevel = true
	// compiler type enforcement
	_ = error(&Error{})
)

// Context provides a way to integrate your own details into the Error.
type Context struct {
	Message  string
	Location Frame
}

func newContext(skip int, message string) Context {
	skip++
	return Context{
		Message:  message,
		Location: CurrentFrame(skip),
	}
}

// Error implements both the Error interface as well as Unwrap.
type Error struct {
	Wraps   error
	Stack   Stack
	Details []Context
	Args    any
	f       Formatter
}

func newError(skip int) *Error {
	skip++
	err := &Error{
		Details: []Context{},
		f:       GetFormatterFunc(),
	}
	if IncludeStack {
		err.Stack = GetStack(skip)
	}
	return err
}

func from(skip int, wraps error) *Error {
	skip++
	if SurfaceLevel {
		err, ok := wraps.(*Error)
		if ok {
			return err
		}
	} else {
		check := &Error{}
		if errors.As(wraps, &check) {
			return check
		}
	}
	err := newError(skip)
	err.Wraps = wraps
	return err
}

// Error implements the error interface.
func (err *Error) Error() string {
	return fmt.Sprintf("%+v", err)
}

// Unwrap allows you to unwrap any internal error which makes the implementation compatible with [errors.As].
func (err *Error) Unwrap() error { return err.Wraps }

func (err *Error) Format(state fmt.State, verb rune) {
	err.f.Format(err, state, verb)
}
