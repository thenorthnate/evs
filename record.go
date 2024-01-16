package evs

import (
	"fmt"
)

// Record is a builder type that is used to build up an [Error]. Once you have created the error
// the way you want it to exist, call [Record.Err] to return an error type.
type Record struct {
	err *Error
}

func newRecord(err *Error) *Record {
	return &Record{
		err: err,
	}
}

// New creates a new [Record] with the given message and the Std error type.
func New(msg string) *Record {
	err := newError(initialSkip)
	detail := newDetail(initialSkip, msg)
	err.Details = append(err.Details, detail)
	return newRecord(err)
}

// Newf creates a new [Record] with the given formatted message and the Std error type.
func Newf(msg string, args ...any) *Record {
	err := newError(initialSkip)
	detail := newDetail(initialSkip, fmt.Sprintf(msg, args...))
	err.Details = append(err.Details, detail)
	return newRecord(err)
}

// From generates a new record from the given error. If the error is nil, the record will contain a nil internal
// [Error]. If the given error is not nil, it first checks to see if it already contains a [Error]. If it does, it
// directly sets the underlying [Error] to that error. Otherwise, it creates a new [Error] that wraps the given error.
// Use this method if you don't intend to "wrap" [Error]s but rather just have one single error that you
// can pass around.
func From(err error) *Record {
	if err == nil {
		return newRecord(nil)
	}
	newErr := from(initialSkip, err)
	return newRecord(newErr)
}

// Msg provides a mechanism to set the error message directly.
func (rec *Record) Msg(msg string) *Record {
	if rec.err == nil {
		return rec
	}
	detail := newDetail(initialSkip, msg)
	rec.err.Details = append(rec.err.Details, detail)
	return rec
}

// Msgf is the same as [Record.Msg] except that it takes a variadic set of arguments.
func (rec *Record) Msgf(msg string, args ...any) *Record {
	if rec.err == nil {
		return rec
	}
	detail := newDetail(initialSkip, fmt.Sprintf(msg, args...))
	rec.err.Details = append(rec.err.Details, detail)
	return rec
}

// DropStack allows you to remove the stacktrace for this specific error. You might use this if you
// generally want the full stacktrace for everything (thus [IncludeStack]==true), but you don't want
// this error to have it.
func (rec *Record) DropStack(message string, args ...any) *Record {
	if rec.err == nil {
		return rec
	}
	rec.err.Stack = Stack{}
	return rec
}

// Set directly assigns the given error to the internal wrapped error. It overrides any previously wrapped
// error that may have already been in place.
func (rec *Record) Set(wraps error) *Record {
	if rec.err == nil {
		return rec
	}
	rec.err.Wraps = wraps
	return rec
}

// Fmt allows you to set the Formatter you'd like to use which dictates how the messages are
// printed out.
func (rec *Record) Fmt(f Formatter) *Record {
	if rec.err == nil {
		return rec
	}
	rec.err.f = f
	return rec
}

// Err returns the error that you've built up via the other methods.
func (rec *Record) Err() error {
	return rec.err
}
