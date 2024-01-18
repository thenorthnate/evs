# evs

[![GoDoc][doc-img]][doc] [![Test][ci-img]][ci]

package `evs` (Error ValueS) is another error generating package. It aims to be simple
and yet contain a useful set of features such as, stack traces, custom formatting, and full
compatibility with the standard library post Go 1.13.

## Example Usage
Create new errors with the `New` function:
```go
func MyFancyFunction() error {
    return evs.New("something went terribly wrong").Err()
}
```

Enhance existing errors (or return nil if there was no error):
```go
func MyFancyFunction() error {
    _, err := fmt.Println("hello, world")
    return evs.From(err).Err()
}
```

## Custom Formatting
You can set your own formatting for these errors! There are two formatters that are built in for 
now but you can easily create your own. The formatters themselves are not exported (trying to keep 
the surface area of the package as small as possible) but functions to set them are exported. If
you want to change formatting for all errors across the board, do something like this:
```go
evs.GetFormatterFunc = evs.JSONFormatter  // by default it is set to the TextFormatter
```

Additionally, you can set the formatter on a per-error basis like this:
```go
err := evs.New("uh oh").Fmt(JSONFormatter()).Err()
```

More generally, you can implement your own by implementing a `evs.Formatter` which will format your
error however you'd like.

[doc-img]: https://pkg.go.dev/badge/github.com/thenorthnate/evs
[doc]: https://pkg.go.dev/github.com/thenorthnate/evs
[ci-img]: https://github.com/thenorthnate/evs/workflows/test/badge.svg
[ci]: https://github.com/thenorthnate/evs/actions
