package evs

import (
	"errors"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	err := New().Msg("terrible error").Err()
	fmt.Println(err)
	require.Contains(t, err.Error(), "*evs.Error[github.com/thenorthnate/evs.Std]: terrible error")
	require.Contains(t, err.Error(), "error_test.go")
}

func TestRecordMsgf(t *testing.T) {
	err := New().Msgf("terrible error: %v", 10).Err()
	fmt.Println(err.Error())
	require.Contains(t, err.Error(), "*evs.Error[github.com/thenorthnate/evs.Std]: terrible error: 10")
}

func TestFrom(t *testing.T) {
	err := errors.New("Hello, world")
	newErr := From(err).Err()
	fmt.Println(newErr.Error())
	require.Contains(t, newErr.Error(), "*evs.Error[github.com/thenorthnate/evs.Std]: Hello, world")
}

func ExampleFrom() {
	err := errors.New("something terrible happened!")
	newErr := From(err).Err()
	if newErr == nil {
		log.Fatal("This should be an error!")
	}
}

func TestErrorKindMatters(t *testing.T) {
	err := New().Msg("terrible error").Err()
	notExpect := &Error[string]{}
	require.False(t, errors.As(err, &notExpect))

	expect := &Error[Std]{}
	require.True(t, errors.As(err, &expect))
	require.Len(t, expect.stack.Frames, 3)
}
