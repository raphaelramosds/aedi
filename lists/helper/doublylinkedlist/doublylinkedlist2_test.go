package helper

import "testing"

func TestGeneral(t *testing.T) {

	var dll DoublyLinkedList

	dll.Init()

	for i := 0; i < 10; i++ {
		dll.Add(i)
	}

	dll.AddOnIndex(12, 3)

	dll.getAll()

}
