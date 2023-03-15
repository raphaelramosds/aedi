package main

import (
	"dca0208.com/lists/helper"
)

func main() {

	var ll helper.LinkedList

	ll.Add(10)
	ll.Add(13)
	ll.Add(9)
	ll.Add(19)

	ll.AddOnIndex(20, 2)
	ll.AddOnIndex(66, 1)
	ll.AddOnIndex(0, 6)

	ll.Set(3, 1)

}
