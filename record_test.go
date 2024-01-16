package evs

import (
	"errors"
	"log"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	err := New("terrible error").Err()
	if !strings.Contains(err.Error(), "*evs.Error: terrible error") {
		t.Fatalf("error \n%v\n did not contain expected output", err.Error())
	}
	if !strings.Contains(err.Error(), "record_test.go") {
		t.Fatal("error did not contain expected output")
	}
}

func TestNewf(t *testing.T) {
	err := Newf("terrible error: %v", 10).Err()
	if !strings.Contains(err.Error(), "*evs.Error: terrible error: 10") {
		t.Fatal("error did not contain expected output")
	}
}

func TestFrom(t *testing.T) {
	err := errors.New("Hello, world")
	newErr := From(err).Err()
	if !strings.Contains(newErr.Error(), "*evs.Error: Hello, world") {
		t.Fatalf("error \n%v\n did not contain expected output", newErr.Error())
	}
}

func TestFromNil(t *testing.T) {
	err := From(nil).Err()
	if err != nil {
		t.Fatal("error was supposed to be nil")
	}
}

func ExampleFrom() {
	err := errors.New("something terrible happened!")
	newErr := From(err).Err()
	if newErr == nil {
		log.Fatal("This should be an error!")
	}
}

func TestRecord_Msg(t *testing.T) {
	err := errors.New("Hello, world")
	newErr := From(err).Msg("got this error").Err()
	if !strings.Contains(newErr.Error(), "got this error") {
		t.Fatalf("error \n%v\n did not contain expected output", newErr.Error())
	}
}

func TestRecord_Msgf(t *testing.T) {
	err := errors.New("Hello, world")
	newErr := From(err).Msgf("got this error %v", 10).Err()
	if !strings.Contains(newErr.Error(), "got this error 10") {
		t.Fatalf("error \n%v\n did not contain expected output", newErr.Error())
	}
}

func TestRecord_DropStack(t *testing.T) {
	err := errors.New("Hello, world")
	newErr := From(err).Msg("got this error").DropStack().Err()
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
	newErr := From(err).Msg("got this error").Set(secondErr).Err()
	if strings.Contains(newErr.Error(), "Hello, world") {
		t.Fatalf("error \n%v\n should not contain that text", newErr.Error())
	}
}
