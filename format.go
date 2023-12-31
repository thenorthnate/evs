package evs

import (
	"fmt"
	"io"
)

var (
	// GetFormatterFunc should return the formatter that gets used in each instantiation of an error. You can
	// supply your own implementation if you would like to change how errors are formatted. See the source
	// code for the [StandardFormatter] to see how it is implemented.
	GetFormatterFunc = DefaultFormatter
)

// Formatter is almost the same as the [fmt.Formatter] but passes the [Error] in as well.
type Formatter interface {
	Format(e *Error, f fmt.State, verb rune)
}

// DefaultFormatter returns the [Formatter] that is used by default. You can use this to restore default
// behavior if you swapped in your own Formatter at some point.
func DefaultFormatter() Formatter {
	return StandardFormatter{}
}

// StandardFormatter is the default [Formatter] used in the errors.
type StandardFormatter struct{}

// Format implements the [Formatter] interface.
func (sf StandardFormatter) Format(e *Error, s fmt.State, verb rune) {
	if e.Wraps != nil {
		formattable, ok := e.Wraps.(fmt.Formatter)
		if ok {
			formattable.Format(s, verb)
		} else {
			_, _ = io.WriteString(s, e.Wraps.Error())
		}
	}

	for _, item := range e.Details {
		item.Format(s, verb)
	}
	e.Stack.Format(s, verb)
}
