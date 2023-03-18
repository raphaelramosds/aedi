package helper

import (
	"testing"
)

func TestGeneral(t *testing.T) {
	var ll LinkedList

	ll.Init()

	for i := 0; i < 10; i++ {
		ll.Add(i)
	}

	ll.RemoveOnIndex(ll.tam - 1)
	ll.Remove()
}
