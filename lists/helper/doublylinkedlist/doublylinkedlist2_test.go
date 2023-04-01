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

func TestAddOnIndexMid(t *testing.T) {
	var list DoublyLinkedList
	//fulfill list with 1's
	for i := 0; i < 10; i++ {
		list.Add(1)
	}

	//add -1, 3 times, on index 2
	for i := 0; i < 3; i++ {
		list.AddOnIndex(-1, 2)
	}

	//check values before index 2 are the same
	for i := 0; i < 8; i++ {
		val := list.Get(i)
		if i >= 2 && i < 5 {
			if val != -1 {
				t.Errorf("%T value on index %d is %d, but we expected it to be -1", list, i, val)
			}
		} else {
			if val != 1 {
				t.Errorf("%T value on index %d is %d, but we expected it to be 1", list, i, val)
			}
		}
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

func TestRemoveOnIndexMid(t *testing.T) {
	var list DoublyLinkedList
	for i := 0; i < 10; i++ {
		list.Add(i)
	}

	//remove on mid
	list.RemoveOnIndex(2)
	for i := 2; i < list.Size(); i++ {
		val := list.Get(i)
		if val != i+1 {
			t.Errorf("%T value on index %d is %d, but we expected it to be %d", list, i, val, i+1)
		}
	}
}

func TestAddAtBegin(t *testing.T) {
	var list DoublyLinkedList

	newElement := 12

	list.Init()

	for i := 0; i < 10; i++ {
		list.Add(i)
	}

	list.AddOnIndex(newElement, 0)

	result := list.Get(0)
	expected := newElement

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

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

func TestSet(t *testing.T) {
	var list DoublyLinkedList

	for i := 0; i < 10; i++ {
		list.Add(1)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			list.Set(-1, i)
		}
	}

	for i := 0; i < 10; i++ {
		val := list.Get(i)
		if i%2 == 0 {
			if val != -1 {
				t.Errorf("%T value on index %d is %d, but we expected it to be -1", list, i, val)
			}
		} else {
			if val == -1 {
				t.Errorf("%T value on index %d is %d, but we expected it to be different from -1", list, i, val)
			}
		}
	}
}
