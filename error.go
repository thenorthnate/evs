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
	Wraps    error
	Stack    Stack
	Messages []string
	Kind     T
	F        Formatter[T]
}

func newError[T any](skip int) *Error[T] {
	skip++
	err := &Error[T]{
		F: DefaultFormatter[T]{},
	}
	if IncludeStack {
		err.Stack = GetStack(skip)
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
	err.Wraps = wraps
	return err
}

// GetKind returns the kind of Error. It is only valid if you set it via a call to [Kind].
func (err *Error[T]) GetKind(kind T) T {
	return err.Kind
}

// Error implements the error interface.
func (err *Error[T]) Error() string {
	return fmt.Sprintf("%+v", err.F)
	// parts := []string{fmt.Sprintf("%T: %v", err, err.message)}
	// if err.wraps != nil {
	// 	if err.message == "" {
	// 		parts[0] = parts[0] + err.wraps.Error()
	// 	} else {
	// 		parts = append(parts, err.wraps.Error())
	// 	}
	// }
	// if len(err.Stack.Frames) > 0 {
	// 	parts = append(parts, "\nWith Stacktrace:", err.Stack.String())
	// }
	// return strings.Join(parts, "\n")
}

// Unwrap allows you to unwrap any internal error which makes the implementation compatible with [errors.As].
func (err *Error[T]) Unwrap() error { return err.Wraps }

// String implements the [fmt.Stringer] interface. It only returns the message from the Error. This is useful
// for structured logging if you don't want to see the stack trace and other details in the message of a log.
// func (err *Error[T]) String() string { return err.Messages }

func (err *Error[T]) Format(f fmt.State, verb rune) {
	err.F.Format(err, f, verb)
}

type DefaultFormatter[T any] struct{}

// Format implements the [Formatter] interface.
func (df DefaultFormatter[T]) Format(e *Error[T], f fmt.State, verb rune) {
	switch verb {
	default:
		fmt.Println("Verb: ", verb)
		_, _ = f.Write([]byte(df.verbose(e)))
	}

}

func (df DefaultFormatter[T]) verbose(e *Error[T]) string {
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
