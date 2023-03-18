package helper

import (
	"testing"
)

func TestRemoveAtBegin(t *testing.T) {
	var list DoublyLinkedList

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
	var list DoublyLinkedList

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

func TestAddAtBegin(t *testing.T) {
	var list DoublyLinkedList

	list.Init()

	for i := 0; i < 10; i++ {
		list.Add(i)
	}

	list.AddOnIndex(12, 0)

	result := list.Get(0)
	expected := 12

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

// In this specific test we're going to add elements at
// doubly linked list, remove one element at begin
// and invert them iterating from tail to head

func TestIterateTailToHead(t *testing.T) {

	var list DoublyLinkedList
	var invList DoublyLinkedList

	expectedArray := [9]int{9, 8, 7, 6, 5, 4, 3, 2, 0}

	invList.Init()
	list.Init()

	for i := 0; i < 10; i++ {
		list.Add(i)
	}

	list.RemoveOnIndex(1)

	aux := list.tail

	for aux != nil {
		invList.Add(aux.val)
		aux = aux.prev
	}

	expected := true

	for i, element := range expectedArray {
		if element != invList.Get(i) {
			expected = false
			break
		}
	}

	if !expected {
		t.Errorf("List isn't inverted")
	}
}
