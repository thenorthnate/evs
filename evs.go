package evs

import "fmt"

const (
	// IOError can be used as the kind for errors related to IO operations.
	IOError ErrorKind = "IOError"
	// TypeError can be used as the kind for errors related to type problems.
	TypeError ErrorKind = "TypeError"
	// ValueError can be used as the kind for errors any time an unexpected value is used.
	ValueError ErrorKind = "ValueError"
)

// ErrorKind defines a category for the error.
type ErrorKind string

type Formatter interface {
	Format(e *Error, f fmt.State, verb rune)
}
