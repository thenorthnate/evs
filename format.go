package evs

import (
	"fmt"
	"io"
	"strings"
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
	return standardFormatter{}
}

// standardFormatter is the default [Formatter] used in the errors.
type standardFormatter struct{}

// Format implements the [Formatter] interface.
func (sf standardFormatter) Format(e *Error, s fmt.State, verb rune) {
	sf.formatWrappedError(e, s, verb)
	sf.formatDetails(e, s, verb)
	sf.formatStack(e.Stack, s, verb)
}

func (sf standardFormatter) formatWrappedError(e *Error, s fmt.State, verb rune) {
	if e.Wraps != nil {
		formattable, ok := e.Wraps.(fmt.Formatter)
		if ok {
			formattable.Format(s, verb)
			_, _ = io.WriteString(s, "\n")
		} else {
			_, _ = io.WriteString(s, e.Wraps.Error()+"\n")
		}
	}
}

func (sf standardFormatter) formatDetails(e *Error, s fmt.State, verb rune) {
	for _, item := range e.Details {
		sf.formatSingleContext(item, s, verb)
	}
}

func (sf standardFormatter) formatSingleContext(ctx Context, s fmt.State, verb rune) {
	sf.formatFrame(ctx.Location, s, verb)
	_, _ = io.WriteString(s, " "+ctx.Message)
	// if ctx.Args != nil {
	// 	formattable, ok := ctx.Args.(fmt.Formatter)
	// 	if ok {
	// 		formattable.Format(s, verb)
	// 	} else {
	// 		_, _ = fmt.Fprintf(s, "\n%v", ctx.Args)
	// 	}

	// }
}

func (sf standardFormatter) formatFrame(frame Frame, s fmt.State, verb rune) {
	fileParts := strings.Split(frame.File, "/")
	switch verb {
	case 's':
		_, _ = fmt.Fprintf(s, "[%v:%v]", fileParts[len(fileParts)-1], frame.Line)
	default:
		_, _ = fmt.Fprintf(s, "%v [%v:%v]", frame.Function, fileParts[len(fileParts)-1], frame.Line)
	}
}

func (sf standardFormatter) formatStack(stack Stack, s fmt.State, verb rune) {
	if len(stack.Frames) == 0 {
		return
	}
	_, _ = io.WriteString(s, "\n\nWith Stacktrace:\n")
	for i, frame := range stack.Frames {
		sf.formatFrame(frame, s, verb)
		if i == len(stack.Frames)-1 {
			break
		}
		_, _ = io.WriteString(s, "\n")
	}
}
