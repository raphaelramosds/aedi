package main

import (
	"dca0208.com/lists/helper"
)

func main() {

	var ll helper.LinkedList

	ll.Add(4)
	ll.Add(12)
	ll.Add(3)
	ll.AddOnIndex(13, 0)
	ll.AddOnIndex(9, 0)
	ll.Remove()
	ll.Remove()
	ll.RemoveOnIndex(0)
	ll.RemoveOnIndex(0)
	ll.RemoveOnIndex(0)
	ll.Add(56)
	ll.Add(24)
	ll.Remove()

}
