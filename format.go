package evs

import (
	"fmt"
	"io"
	"strings"
)

var (
	// GetFormatterFunc should return the formatter that gets used in each instantiation of an error. You can
	// supply your own implementation if you would like to change how errors are formatted. See the source
	// code for the [textFormatter] to see how it is implemented.
	GetFormatterFunc = TextFormatter
)

// Formatter is almost the same as the [fmt.Formatter] but passes the [Error] in as well.
type Formatter interface {
	Format(e *Error, f fmt.State, verb rune)
}

// TextFormatter returns the [Formatter] that is used by default. You can use this to restore default
// behavior if you swapped in your own Formatter at some point.
func TextFormatter() Formatter {
	return textFormatter{}
}

// textFormatter is the default [Formatter] used in the errors.
type textFormatter struct{}

// Format implements the [Formatter] interface.
func (f textFormatter) Format(e *Error, s fmt.State, verb rune) {
	f.formatWrappedError(e, s, verb)
	f.formatDetails(e, s, verb)
	f.formatStack(e.Stack, s, verb)
}

func (f textFormatter) formatWrappedError(e *Error, s fmt.State, verb rune) {
	if e.Wraps != nil {
		formattable, ok := e.Wraps.(fmt.Formatter)
		if ok {
			formattable.Format(s, verb)
			_, _ = io.WriteString(s, "\n")
		} else {
			_, _ = fmt.Fprintf(s, "%s\n", e.Wraps.Error())
		}
	}
}

func (f textFormatter) formatDetails(e *Error, s fmt.State, verb rune) {
	_, _ = fmt.Fprintf(s, "%v", e.Details)
}

func (f textFormatter) formatFrame(frame Frame, s fmt.State, verb rune) {
	fileParts := strings.Split(frame.File, "/")
	switch verb {
	case 's':
		_, _ = fmt.Fprintf(s, "[%v:%v]", fileParts[len(fileParts)-1], frame.Line)
	default:
		_, _ = fmt.Fprintf(s, "%v [%v:%v]", frame.Function, fileParts[len(fileParts)-1], frame.Line)
	}
}

func (f textFormatter) formatStack(stack Stack, s fmt.State, verb rune) {
	if len(stack.Frames) == 0 {
		return
	}
	_, _ = io.WriteString(s, "\n\nWith Stacktrace:\n")
	for i, frame := range stack.Frames {
		f.formatFrame(frame, s, verb)
		if i == len(stack.Frames)-1 {
			break
		}
		_, _ = io.WriteString(s, "\n")
	}
}
