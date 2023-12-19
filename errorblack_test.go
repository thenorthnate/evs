package evs_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thenorthnate/evs"
)

func TestFrom(t *testing.T) {
	err := errors.New("Hello, world")
	newErr := evs.From(err).Err()
	require.Contains(t, newErr.Error(), "*evs.Error[github.com/thenorthnate/evs.Std]: Hello, world")
}
