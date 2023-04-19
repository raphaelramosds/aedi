package main

import (
	"dca0208.com/sortings/helper"
)

func main() {
	s := helper.Init(10)
	helper.Display(s[:])
	helper.SelectionSort(s[:])
	helper.Display(s[:])
}
