package main

import (
	"fmt"

	"dca0208.com/bst/helper"
)

func main() {

	// List of prime numbers
	primes := [8]int{3, 1, 5, 7, 11, 13, 3, 5}
	root := &helper.BSTNode{}

	// Inserting elements
	for _, element := range primes {
		root.Insert(element)
	}

	// Traversals

	// root.PrintIn()
	// root.PrintPre()
	// root.PrintPos()

	// Searches
	fmt.Println(root.Search(3))

	// Max and minimum
	fmt.Println(root.Max())
	fmt.Println(root.Min())

}
