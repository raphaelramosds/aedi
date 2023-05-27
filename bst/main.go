package main

import "dca0208.com/bst/helper"

func main() {
	// Silence is golden
	root := &helper.BSTNode{}

	root.Insert(3)
	root.Insert(5)
	root.Insert(7)
	root.Insert(11)

	root.PrintIn()

	root = root.Remove(3)

	root.PrintIn()
}
