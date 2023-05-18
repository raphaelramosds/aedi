package main

import (
	"fmt"

	"dca0208.com/bst/helper"
)

func main() {
	primes := [6]int{3, 1, 5, 7, 11, 13}
	bst := helper.BinarySearchTree{}

	for _, element := range primes {
		bst.Insert(element)
	}

	if bst.Search(5) {
		fmt.Println("Found it")
	} else {
		fmt.Println("Not found")
	}

}
