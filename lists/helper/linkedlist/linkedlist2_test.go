package helper

import (
	"testing"
)

func TestInit(t *testing.T) {
	var ll LinkedList

	ll.Init()

	result := ll.tam
	expected := 0

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

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
