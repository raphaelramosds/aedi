package helper

import (
	"testing"
)

func TestAddIntoAnEmptyList(t *testing.T) {
	var ll LinkedList

	newNode := 1

	ll.Add(newNode)

	result := ll.Get(0)
	expected := newNode

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}

}
