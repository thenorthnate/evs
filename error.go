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

func newError[T any](skip int, message string) *Error[T] {
	skip++
	err := &Error[T]{
		message: message,
	}
	if IncludeStack {
		err.stack = GetStack(skip)
	}
	return err
}

// New creates a new [Error] with the given message and the Std error type.
func New(message string) *Error[Std] {
	return newError[Std](initialSkip, message)
}

// NewT creates a new [Error] with the given message and the generic type specified.
func NewT[T any](message string) *Error[T] {
	return newError[T](initialSkip, message)
}

// Newf creates a new [Error] with the given message formatted with the given arguments and the Std error type.
func Newf(message string, args ...any) *Error[Std] {
	return newError[Std](initialSkip, fmt.Sprintf(message, args...))
}

// NewTf creates a new [Error] with the given message formatted with the given arguments.
func NewTf[T any](message string, args ...any) *Error[T] {
	return newError[T](initialSkip, fmt.Sprintf(message, args...))
}

// From first checks the given error to see if it already contains a [*Error]. If it does, it directly
// returns the underlying [*Error]. Otherwise, it creates a new [*Error] that wraps the given error.
// Use this method if you don't intend to "wrap" [*Error]'s but rather just have one single error that you
// can pass around. Do note that this method assumes errors with type [*Error[Std]]. If you generated
// the error with a generic type, use [FromT] instead.
func From(wraps error) *Error[Std] {
	check := &Error[Std]{}
	if errors.As(wraps, &check) {
		return check
	}
	return newError[Std](initialSkip, "").Set(wraps)
}

// FromT works exactly the same way as [From] but is a generic method that requires you provide the generic
// type. This way you can check for errors that use your own generic type such as those created by [NewT].
func FromT[T any](wraps error) *Error[T] {
	check := &Error[T]{}
	if errors.As(wraps, &check) {
		return check
	}
	return newError[T](initialSkip, "").Set(wraps)
}

// Set directly assigns the given error to the internal wrapped error. It overrides any previously wrapped
// error that may have already been in place.
func (err *Error[T]) Set(wraps error) *Error[T] {
	err.wraps = wraps
	return err
}

// Msg provides a mechanism to set the error message directly.
func (err *Error[T]) Msg(message string) *Error[T] {
	err.message = message
	return err
}

// Msgf is the same as [Error.Msg] except that it takes a variadic set of arguments.
func (err *Error[T]) Msgf(message string, args ...any) *Error[T] {
	err.message = fmt.Sprintf(message, args...)
	return err
}

// DropStack allows you to remove the stacktrace for this specific error. You might use this if you
// generally want the full stacktrace for everything (thus [IncludeStack]==true), but you don't want
// this error to have it.
func (err *Error[T]) DropStack(message string, args ...any) *Error[T] {
	err.stack = Stack{}
	return err
}

// Kind allows you to set the error kind. This is not used internally, but can be useful for your
// own use cases if you want to set a specific (computer-parsable) reason for the error.
func (err *Error[T]) Kind(kind T) *Error[T] {
	err.kind = kind
	return err
}

// GetKind returns the kind of Error. It is only valid if you set it via a call to [Kind].
func (err *Error[T]) GetKind(kind T) T {
	return err.kind
}

// Error implements the error interface.
func (err *Error[T]) Error() string {
	parts := []string{fmt.Sprintf("%T: %v", *new(T), err.message)}
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
