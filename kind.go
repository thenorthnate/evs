package evs

const (
	// IOError can be used as the kind for errors related to IO operations.
	IOError Std = "IOError"
	// TypeError can be used as the kind for errors related to type problems.
	TypeError Std = "TypeError"
	// ValueError can be used as the kind for errors any time an unexpected value is used.
	ValueError Std = "ValueError"
)

// Std is the standard type used to generate new Errors. It is simply an alias for a string.
// You can use your own types if desired.
type Std string
