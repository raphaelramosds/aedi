package helper

import "testing"

func TestAddManyValues(t *testing.T) {
	var dll DoublyLinkedList

	dll.Init()

	dll.Add(9)
	dll.Add(5)
	dll.Add(4)
	dll.Add(2)

	expected := 5
	result := dll.Get(1)

	if expected != result {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
