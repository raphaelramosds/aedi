package helper

import (
	"testing"
)

func TestRemoveAtBegin(t *testing.T) {
	var list LinkedList

	list.Init()

	for i := 0; i < 10; i++ {
		list.Add(i)
	}

	list.RemoveOnIndex(0)

	result := list.Get(0)
	expected := 1

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestRemoveAtEnd(t *testing.T) {
	var list LinkedList

	list.Init()

	for i := 0; i < 10; i++ {
		list.Add(i)
	}

	list.RemoveOnIndex(list.tam - 1)

	result := list.Get(list.tam - 1)
	expected := 8

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestRemoveAtASpecificIndex(t *testing.T) {
	var list LinkedList

	index := 2

	list.Init()

	for i := 0; i < 10; i++ {
		list.Add(i)
	}

	neighbor := list.Get(index + 1)
	list.RemoveOnIndex(index)

	result := list.Get(index)
	expected := neighbor

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestAddAtBegin(t *testing.T) {
	var list LinkedList

	newElement := 5
	index := 0

	list.Init()

	for i := 0; i < 10; i++ {
		list.Add(i)
	}

	list.AddOnIndex(newElement, index)

	result := list.Get(index)
	expected := newElement

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestAddAtASpecificIndex(t *testing.T) {
	var list LinkedList

	newElement := 5
	index := 2

	list.Init()

	for i := 0; i < 10; i++ {
		list.Add(i)
	}

	list.AddOnIndex(newElement, index)

	result := list.Get(index)
	expected := newElement

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestAddIntoAnEmptyList(t *testing.T) {
	var list LinkedList

	newElement := 1

	list.Init()
	list.Add(newElement)

	result := list.Get(0)
	expected := newElement

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestAddIntoAFullFilledList(t *testing.T) {

	var list LinkedList
	newElement := 4

	list.Init()

	for i := 0; i < 10; i++ {
		list.Add(i)
	}

	list.Add(newElement)

	result := list.Get(10)
	expected := newElement

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
