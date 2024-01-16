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

func TestRecordMsgf(t *testing.T) {
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

func ExampleFrom() {
	err := errors.New("something terrible happened!")
	newErr := From(err).Err()
	if newErr == nil {
		log.Fatal("This should be an error!")
	}
}
