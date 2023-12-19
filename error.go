package evs

import (
	"errors"
	"fmt"
	"strings"
)

const (
	initialSkip = 1
)

var (
	// IncludeStack is used to determine whether or not a stacktrace should be captured with
	// new errors. By default it is set to true.
	IncludeStack = true
	// compiler type enforcement
	_ = error(&Error[Std]{})
)

// Error implements both the Error interface as well as Unwrap.
type Error[T any] struct {
	wraps   error
	stack   Stack
	message string
	kind    T
}

func newError[T any](skip int) *Error[T] {
	skip++
	err := &Error[T]{}
	if IncludeStack {
		err.stack = GetStack(skip)
	}
	return err
}

func from[T any](skip int, wraps error) *Error[T] {
	skip++
	check := &Error[T]{}
	if errors.As(wraps, &check) {
		return check
	}
	err := newError[T](skip)
	err.wraps = wraps
	return err
}

// GetKind returns the kind of Error. It is only valid if you set it via a call to [Kind].
func (err *Error[T]) GetKind(kind T) T {
	return err.kind
}

// Error implements the error interface.
func (err *Error[T]) Error() string {
	parts := []string{fmt.Sprintf("%T: %v", err, err.message)}
	if err.wraps != nil {
		if err.message == "" {
			parts[0] = parts[0] + err.wraps.Error()
		} else {
			parts = append(parts, err.wraps.Error())
		}
	}
	if len(err.stack.Frames) > 0 {
		parts = append(parts, "\nWith Stacktrace:", err.stack.String())
	}
	return strings.Join(parts, "\n")
}

// Unwrap allows you to unwrap any internal error which makes the implementation compatible with [errors.As].
func (err *Error[T]) Unwrap() error { return err.wraps }

// String implements the [fmt.Stringer] interface. It only returns the message from the Error. This is useful
// for structured logging if you don't want to see the stack trace and other details in the message of a log.
func (err *Error[T]) String() string { return err.message }
