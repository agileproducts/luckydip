package main

import (
	"reflect"
	"testing"
)

func TestPick(t *testing.T) {
	actual := Pick(3)
	expected := []string{"wot","not","dot"}

	expectedType := reflect.TypeOf(expected)
	actualType := reflect.TypeOf(actual)

	expectedLength := 3
	actualLength := len(actual)

	// it should return a slice of strings
	if actualType != expectedType {
		t.Errorf("expected %v got %v", expectedType, actualType)
	}

	// it should return the number of strings requested
	if actualLength != expectedLength {
		t.Errorf("expected slice of length %v, got %v", expectedLength,actualLength)
	}
}

/* needs to
 - return random selection of required length from pool
 - not return the same element twice within that selection
 - if request selection bigger than pool then error or cap out
 - load pool from file database
 - have a web and a command line interface
 - replace dummy strings with real article types
*/