package helper

import (
	"testing"
)

func TestInit(t *testing.T) {
	var al ArrayList
	al.Init()
	if len(al.values) == 0 {
		t.Errorf("Array list wasn't created")
	}
}

func TestAddIntoAnEmptyList(t *testing.T) {
	var al ArrayList

	newElement := 1

	al.Init()
	al.Add(newElement)

	result := al.values[0]
	expected := newElement

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestAddIntoAFilledList(t *testing.T) {
	var al ArrayList

	al.Init()

	for i := 0; i < len(al.values); i++ {
		al.Add(12)
	}

	al.Add(4)

	result := len(al.values)
	expected := 20

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
