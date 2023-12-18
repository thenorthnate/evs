# evs

package `evs` (error values with sugar) is another error generating package. It aims to be simple
and yet contain a useful set of features such as generic error types, stack traces, and full
compatibility with the standard library post Go 1.13.

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
