package evs

import (
	"fmt"
)

// Record is a builder type that is used to build up an [Error]. Once you have created the error
// the way you want it to exist, call [Record.Err] to return an error type.
type Record[T any] struct {
	err *Error[T]
}

func newRecord[T any](err *Error[T]) *Record[T] {
	return &Record[T]{
		err: err,
	}
}

// New creates a new [Record] with the given message and the Std error type.
func New(msg string) *Record[Std] {
	err := newError[Std](initialSkip)
	err.message = msg
	return newRecord(err)
}

// Newf creates a new [Record] with the given formatted message and the Std error type.
func Newf(msg string, args ...any) *Record[Std] {
	err := newError[Std](initialSkip)
	err.message = fmt.Sprintf(msg, args...)
	return newRecord(err)
}

// NewT creates a new [Record] with the given message and the generic type specified.
func NewT[T any](msg string) *Record[T] {
	err := newError[T](initialSkip)
	err.message = msg
	return newRecord(err)
}

// NewTf creates a new [Record] with the given formatted message and the generic type specified.
func NewTf[T any](msg string, args ...any) *Record[T] {
	err := newError[T](initialSkip)
	err.message = fmt.Sprintf(msg, args...)
	return newRecord(err)
}

// From generates a new record from the given error. If the error is nil, the record will contain a nil internal
// [Error]. If the given error is not nil, it first checks to see if it already contains a [Error]. If it does, it
// directly sets the underlying [Error] to that error. Otherwise, it creates a new [Error] that wraps the given error.
// Use this method if you don't intend to "wrap" [Error]s but rather just have one single error that you
// can pass around.
func From(err error) *Record[Std] {
	if err == nil {
		return newRecord[Std](nil)
	}
	newErr := from[Std](initialSkip, err)
	return newRecord(newErr)
}

// FromT is exactly like [From] except that it requires you to specify the generic type to use instead of defaulting
// to the [Std] type.
func FromT[T any](err error) *Record[T] {
	if err == nil {
		return newRecord[T](nil)
	}
	newErr := from[T](initialSkip, err)
	return newRecord(newErr)
}

// Msg provides a mechanism to set the error message directly.
func (rec *Record[T]) Msg(message string) *Record[T] {
	if rec.err == nil {
		return rec
	}
	rec.err.message = message
	return rec
}

// Msgf is the same as [Record.Msg] except that it takes a variadic set of arguments.
func (rec *Record[T]) Msgf(message string, args ...any) *Record[T] {
	if rec.err == nil {
		return rec
	}
	rec.err.message = fmt.Sprintf(message, args...)
	return rec
}

// DropStack allows you to remove the stacktrace for this specific error. You might use this if you
// generally want the full stacktrace for everything (thus [IncludeStack]==true), but you don't want
// this error to have it.
func (rec *Record[T]) DropStack(message string, args ...any) *Record[T] {
	if rec.err == nil {
		return rec
	}
	rec.err.stack = Stack{}
	return rec
}

// Kind allows you to set the error kind. This is not used internally, but can be useful for your
// own use cases if you want to set a specific (computer-parsable) reason for the error.
func (rec *Record[T]) Kind(kind T) *Record[T] {
	if rec.err == nil {
		return rec
	}
	rec.err.kind = kind
	return rec
}

// Set directly assigns the given error to the internal wrapped error. It overrides any previously wrapped
// error that may have already been in place.
func (rec *Record[T]) Set(wraps error) *Record[T] {
	if rec.err == nil {
		return rec
	}
	rec.err.wraps = wraps
	return rec
}

// Err returns the error that you've built up via the other methods.
func (rec *Record[T]) Err() error {
	return rec.err
}
