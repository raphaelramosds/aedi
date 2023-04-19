package main

import (
	"fmt"

	"dca0208.com/search/helper"
)

func main() {
	// Silence is golden
	list := [5]int{1, 2, 3, 4, 5}
	fmt.Println(helper.BSearch(0, len(list), list, 5))
}
