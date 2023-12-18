package evs

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	err := New("terrible error")
	require.Contains(t, err.Error(), "evs.Std: terrible error")
	require.Contains(t, err.Error(), "[error_test.go:11]")
}

func TestNewf(t *testing.T) {
	err := Newf("terrible error: %v", 10)
	require.Contains(t, err.Error(), "evs.Std: terrible error: 10")
}

func TestFrom(t *testing.T) {
	err := errors.New("Hello, world")
	newErr := From(err)
	require.Contains(t, newErr.Error(), "evs.Std: Hello, world")
}

func TestErrorKindMatters(t *testing.T) {
	err := New("terrible error")
	notExpect := &Error[string]{}
	require.False(t, errors.As(err, &notExpect))

	expect := &Error[Std]{}
	require.True(t, errors.As(err, &expect))
	require.Len(t, expect.stack.Frames, 3)
}

func TestUnwrap(t *testing.T) {
	err1 := New("terrible error")
	err2 := New("second error").Set(err1)

	err := errors.Unwrap(err2)
	err1unwrapped, ok := err.(*Error[Std])
	require.True(t, ok)
	require.EqualValues(t, err1, err1unwrapped)
}
