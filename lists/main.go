package main

import (
	"dca0208.com/lists/helper"
)

func main() {

	var ll helper.LinkedList

	ll.Add(4)
	ll.Add(5)
	ll.Add(2)
	ll.RemoveOnIndex(2)
	ll.RemoveOnIndex(1)
	ll.Add(9)
	ll.Add(3)
	ll.Add(7)
	ll.Add(12)
	ll.Add(29)
	ll.Add(90)
	ll.Remove()
	ll.AddOnIndex(19, 4)
	ll.RemoveOnIndex(3)
	ll.Remove()
	ll.Remove()
	ll.Remove()
	ll.Remove()

}
