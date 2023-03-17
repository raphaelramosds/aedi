package helper

import (
	"testing"
)

func TestGeneral(t *testing.T) {

	var dll DoublyLinkedList

	dll.Init()

	for i := 0; i < 10; i++ {
		dll.Add(i)
	}

	dll.AddOnIndex(12, dll.tam-1)
	dll.Add(98)
	dll.Add(50)
	dll.AddOnIndex(9, 0)

	dll.getAll()

}
