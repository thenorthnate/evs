package evs

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

const (
	initialSkip = 1
)

var (
	// IncludeStack is used to determine whether or not a stacktrace should be captured with
	// new errors. By default it is set to true.
	IncludeStack               = true
	DefaultFormatter Formatter = StandardFormatter{}
	// compiler type enforcement
	_ = error(&Error{})
)

// Error implements both the Error interface as well as Unwrap.
type Error struct {
	Wraps    error
	Stack    Stack
	Messages []string
	Kind     ErrorKind
	f        Formatter
}

func newError(skip int) *Error {
	skip++
	err := &Error{}
	if IncludeStack {
		err.Stack = GetStack(skip)
	}
	return err
}

func from(skip int, wraps error) *Error {
	skip++
	check := &Error{}
	if errors.As(wraps, &check) {
		return check
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

// String implements the [fmt.Stringer] interface. It only returns the message from the Error. This is useful
// for structured logging if you don't want to see the stack trace and other details in the message of a log.
// func (err *Error[T]) String() string { return err.Messages }

func (err *Error) Format(state fmt.State, verb rune) {
	if err.f != nil {
		err.f.Format(err, state, verb)
	} else {
		DefaultFormatter.Format(err, state, verb)
	}
}

type StandardFormatter struct{}

// Format implements the [Formatter] interface.
func (sf StandardFormatter) Format(e *Error, s fmt.State, verb rune) {
	switch verb {
	case 'v':
		switch {
		case s.Flag('+'):
			// do nothing for now...
		}
		// do nothing for now...
		fallthrough
	case 's':
		// do nothing for now...
		fallthrough
	default:
		_, _ = io.WriteString(s, sf.verbose(e))
	}

}

func (sf StandardFormatter) verbose(e *Error) string {
	parts := []string{fmt.Sprintf("%T: %v", e, strings.Join(e.Messages, "\n"))}
	if e.Wraps != nil {
		if len(e.Messages) > 0 {
			parts[0] = parts[0] + e.Wraps.Error()
		} else {
			parts = append(parts, e.Wraps.Error())
		}
	}
	if len(e.Stack.Frames) > 0 {
		parts = append(parts, "\nWith Stacktrace:", e.Stack.String())
	}
	return strings.Join(parts, "\n")
}
