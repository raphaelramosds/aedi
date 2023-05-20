package main

import "dca0208.com/bst/helper"

func main() {
	primes := [8]int{3, 1, 5, 7, 11, 13, 3, 5}
	root := &helper.BSTNode{}

	// Inserting elements
	for _, element := range primes {
		root.Insert(element)
	}

	// Traversals
	// root.PrintIn()
	root.PrintPre()

}
