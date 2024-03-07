package evs

import (
	"errors"
	"fmt"
)

const (
	KindIO      = "IO"
	KindType    = "Type"
	KindUnknown = ""
	KindValue   = "Value"
	initialSkip = 1
)

var (
	// IncludeStack is used to determine whether or not a stacktrace should be captured with
	// new errors. By default it is set to true.
	IncludeStack = true
	// InspectFull controls how [From] operates. By default, the full error stack will be inspected
	// via [errors.As]. If any [evs.Error] exists within the stack, that error is extracted and returned.
	// You can turn this behavior off, by setting InspectFull to false. This will then only check the
	// error itself (without calling unwrap).
	InspectFull = true
	// compiler type enforcement
	_ = error(&Error{})
)

type Kind string

// Error implements both the Error interface as well as Unwrap.
type Error struct {
	Wraps   error
	Stack   Stack
	Details []string
	Kind    Kind
	f       Formatter
}

func newError(skip int) *Error {
	skip++
	err := &Error{
		f: GetFormatterFunc(),
	}
	if IncludeStack {
		err.Stack = GetStack(skip)
	}
	return err
}

func from(skip int, wraps error) *Error {
	skip++
	if InspectFull {
		check := &Error{}
		if errors.As(wraps, &check) {
			return check
		}
	} else {
		err, ok := wraps.(*Error)
		if ok {
			return err
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

// Format implements the [fmt.Formatter] interface.
func (err *Error) Format(state fmt.State, verb rune) {
	err.f.Format(err, state, verb)
}

// KindOf returns the Kind of this error. If it cannot determine the Kind (e.g. because
// maybe the provided error is not an [evs.Error]) it returns [KindUnknown].
func KindOf(err error) Kind {
	e := &Error{}
	if errors.As(err, &e) {
		return e.Kind
	}
	return KindUnknown
}
