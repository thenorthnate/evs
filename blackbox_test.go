package evs_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/thenorthnate/evs"
)

func TestBlackBoxFrom(t *testing.T) {
	err := errors.New("Hello, world")
	newErr := evs.From(err).Err()
	if !strings.Contains(newErr.Error(), "*evs.Error[github.com/thenorthnate/evs.Std]: Hello, world") {
		t.Fatal("error did not contain expected output")
	}
}
