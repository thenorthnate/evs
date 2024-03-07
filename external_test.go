package evs_test

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/thenorthnate/evs"
)

func ExampleNew() {
	err := evs.New("something terrible happened!").Err()
	if err == nil {
		log.Fatal("This should be an error!")
	}
}

func ExampleFrom() {
	err := errors.New("something terrible happened!")
	newErr := evs.From(err).Err()
	if newErr == nil {
		log.Fatal("This should be an error!")
	}
}

func TestNew(t *testing.T) {
	err := evs.New("terrible error").Err()
	if !strings.Contains(err.Error(), "*evs.Error: terrible error") {
		t.Fatalf("error \n%v\n did not contain expected output", err.Error())
	}
	if !strings.Contains(err.Error(), "external_test.go") {
		t.Fatal("error did not contain expected output")
	}
}

func TestNewf(t *testing.T) {
	err := evs.Newf("terrible error: %v", 10).Err()
	if !strings.Contains(err.Error(), "*evs.Error: terrible error: 10") {
		t.Fatal("error did not contain expected output")
	}
}

func TestFrom(t *testing.T) {
	err := errors.New("Hello, world")
	newErr := evs.From(err).Err()
	if !strings.Contains(newErr.Error(), "*evs.Error: Hello, world") {
		t.Fatalf("error \n%v\n did not contain expected output", newErr.Error())
	}
}

func TestFromNil(t *testing.T) {
	err := evs.From(nil).Err()
	if err != nil {
		t.Fatal("error was supposed to be nil")
	}
}

func TestUnwrap(t *testing.T) {
	err := evs.From(errors.New("uh oh")).Err()
	evsErr, ok := err.(*evs.Error)
	if !ok {
		t.Fatalf("expected *evs.Error but got %T", err)
	}
	subErr := evsErr.Unwrap()
	if subErr.Error() != "uh oh" {
		t.Fatalf("subErr has invalid contents: %v", subErr.Error())
	}
}

func TestRecord_Msg(t *testing.T) {
	err := errors.New("Hello, world")
	newErr := evs.From(err).Msg("got this error").Err()
	if !strings.Contains(newErr.Error(), "got this error") {
		t.Fatalf("error \n%v\n did not contain expected output", newErr.Error())
	}
}

func TestRecord_Msgf(t *testing.T) {
	err := errors.New("Hello, world")
	newErr := evs.From(err).Msgf("got this error %v", 10).Err()
	if !strings.Contains(newErr.Error(), "got this error 10") {
		t.Fatalf("error \n%v\n did not contain expected output", newErr.Error())
	}
}

func TestRecord_DropStack(t *testing.T) {
	err := errors.New("Hello, world")
	newErr := evs.From(err).Msg("got this error").DropStack().Err()
	if !strings.Contains(newErr.Error(), "got this error") {
		t.Fatalf("error \n%v\n did not contain expected output", newErr.Error())
	}
	if strings.Contains(newErr.Error(), "With Stacktrace") {
		t.Fatalf("error \n%v\n should not have a stacktrace", newErr.Error())
	}
}

func TestRecord_Set(t *testing.T) {
	err := errors.New("Hello, world")
	secondErr := errors.New("second error")
	newErr := evs.From(err).Msg("got this error").Set(secondErr).Err()
	if strings.Contains(newErr.Error(), "Hello, world") {
		t.Fatalf("error \n%v\n should not contain that text", newErr.Error())
	}
}

func TestFromExisting_InspectFull(t *testing.T) {
	first := evs.New("bad day").Err()
	second := evs.From(first).Err()
	if second == nil {
		t.Fatal("error should not be nil")
	}
	expect := &evs.Error{}
	if !errors.As(second, &expect) {
		t.Fatal("expected error to have type Error but it did not")
	}
	if expect.Wraps != nil {
		t.Fatalf("wraps should be nil but found: %v", expect.Wraps)
	}
}

func TestFromExisting_InspectFullFalse(t *testing.T) {
	evs.InspectFull = false
	defer func() {
		evs.InspectFull = true
	}()
	first := evs.New("bad day").Err()
	second := evs.From(first).Err()
	if second == nil {
		t.Fatal("error should not be nil")
	}
	expect := &evs.Error{}
	if !errors.As(second, &expect) {
		t.Fatal("expected error to have type Error but it did not")
	}
	if expect.Wraps != nil {
		t.Fatalf("wraps should be nil but found: %v", expect.Wraps)
	}
}

func TestFrom_InspectFullTrue(t *testing.T) {
	first := evs.New("bad day").Err()
	second := fmt.Errorf("bad error: %w", first)
	third := evs.From(second).Err()
	if third == nil {
		t.Fatal("error should not be nil")
	}
	expect := &evs.Error{}
	if !errors.As(third, &expect) {
		t.Fatal("expected error to have type Error but it did not")
	}
	if expect.Wraps != nil {
		t.Fatalf("wraps should be nil but found: %v", expect.Wraps)
	}
}

func TestFrom_InspectFullFalse(t *testing.T) {
	evs.InspectFull = false
	defer func() {
		evs.InspectFull = true
	}()
	first := evs.New("bad day").Err()
	second := fmt.Errorf("bad error: %w", first)
	third := evs.From(second).Err()
	if third == nil {
		t.Fatal("error should not be nil")
	}
	expect, ok := third.(*evs.Error)
	if !ok {
		t.Fatalf("expected type *Error, but got %T", third)
	}
	if expect.Wraps == nil {
		t.Fatal("wraps should not be nil but it was")
	}
}

func TestKindOf(t *testing.T) {
	first := evs.New("bad day").Err()
	k := evs.KindOf(first)
	if k != evs.KindUnknown {
		t.Fatal("kind was supposed to be unknown")
	}
}

func TestKindOf_OtherError(t *testing.T) {
	first := errors.New("uh oh!")
	k := evs.KindOf(first)
	if k != evs.KindUnknown {
		t.Fatal("kind was supposed to be unknown")
	}
}

func TestKind(t *testing.T) {
	first := evs.New("bad day").Kind(evs.KindIO).Err()
	k := evs.KindOf(first)
	if k != evs.KindIO {
		t.Fatal("kind was supposed to be IO")
	}
}

func TestKind_FromNil(t *testing.T) {
	err := evs.From(nil).Kind(evs.KindIO).Err()
	if err != nil {
		t.Fatal("error was supposed to be nil")
	}
}
