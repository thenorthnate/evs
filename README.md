# evs

[![GoDoc][doc-img]][doc] [![Test][ci-img]][ci]

package `evs` (error values with sugar) is another error generating package. It aims to be simple
and yet contain a useful set of features such as generic error types, stack traces, and full
compatibility with the standard library post Go 1.13.

- eva (Error VAlues)
- eye (Expand Your Errors)
- rev (Review Error Values)

## Example Usage
Create new errors with the `New` function (or specify your own generic type with `NewT`):
```go
func MyFancyFunction() error {
    return evs.New("something went terribly wrong")
}
```

Enhance existing errors (or return nil if there was no error):
```go
func MyFancyFunction() error {
    _, err := fmt.Println("hello, world")
    return evs.From(err)
}
```

[doc-img]: https://pkg.go.dev/badge/github.com/thenorthnate/evs
[doc]: https://pkg.go.dev/github.com/thenorthnate/evs
[ci-img]: https://github.com/thenorthnate/evs/workflows/test/badge.svg
[ci]: https://github.com/thenorthnate/evs/actions
