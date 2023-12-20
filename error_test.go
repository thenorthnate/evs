//go:build whitebox

package evs

import (
	"errors"
	"log"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	err := New("terrible error").Err()
	if !strings.Contains(err.Error(), "*evs.Error[github.com/thenorthnate/evs.Std]: terrible error") {
		t.Fatal("error did not contain expected output")
	}
	if !strings.Contains(err.Error(), "error_test.go") {
		t.Fatal("error did not contain expected output")
	}
}

func TestRecordMsgf(t *testing.T) {
	err := Newf("terrible error: %v", 10).Err()
	if !strings.Contains(err.Error(), "*evs.Error[github.com/thenorthnate/evs.Std]: terrible error: 10") {
		t.Fatal("error did not contain expected output")
	}
}

func TestFrom(t *testing.T) {
	err := errors.New("Hello, world")
	newErr := From(err).Err()
	if !strings.Contains(newErr.Error(), "*evs.Error[github.com/thenorthnate/evs.Std]: Hello, world") {
		t.Fatal("error did not contain expected output")
	}
}

func ExampleFrom() {
	err := errors.New("something terrible happened!")
	newErr := From(err).Err()
	if newErr == nil {
		log.Fatal("This should be an error!")
	}
}

func TestErrorKindMatters(t *testing.T) {
	err := New("terrible error").Err()
	notExpect := &Error[string]{}
	if errors.As(err, &notExpect) {
		t.Fatalf("did not expect %T as the error type", notExpect)
	}

	expect := &Error[Std]{}
	if !errors.As(err, &expect) {
		t.Fatalf("expected %T but that was not the type", expect)
	}
	if len(expect.stack.Frames) != 3 {
		t.Fatalf("expected there to be 3 stack frames but got %v", len(expect.stack.Frames))
	}
}
