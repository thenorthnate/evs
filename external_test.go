package evs_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/thenorthnate/evs"
)

func TestExternalFrom(t *testing.T) {
	err := errors.New("Hello, world")
	newErr := evs.From(err).Err()
	if !strings.Contains(newErr.Error(), "*evs.Error[github.com/thenorthnate/evs.Std]: Hello, world") {
		t.Fatalf("%v does not contain expected content", newErr.Error())
	}
}
