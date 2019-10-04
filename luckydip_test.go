package luckydip

import (
	"reflect"
	"testing"
)

func TestRandomSubset(t *testing.T) {
	set := []string{"one","two","three","four","five"}

	actual := RandomSubset(set,3)
	actualType := reflect.TypeOf(actual)
	actualLength := len(actual)

	expectedType := reflect.TypeOf(set)
	expectedLength := 3

	//it should return a slice of the same type that was supplied
	if actualType != expectedType {
		t.Errorf("expected %v, got %v", expectedType, actualType)
	}

	//it should return a slice of the requested length
	if actualLength != expectedLength {
		t.Errorf("expected %v, got %v", expectedLength, actualLength)
	}
}



/* needs to
 - return random selection of required length from pool
 - not return the same element twice within that selection. 
 -  Might need to make it a map rather than a slice cos removing elements will change index?
 - https://yourbasic.org/golang/delete-element-slice/
 - if request selection bigger than pool then error or cap out
 - load pool from file database
 - have a web and a command line interface
 - replace dummy strings with real article types
*/