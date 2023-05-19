package main

import (
	"fmt"
	"time"

	"dca0208.com/sortings/helper"
)

func main() {
	n := 200000
	v := make([]int, n)

	for i := 0; i < n; i++ {
		v[i] = n - i
	}

	start := time.Now()
	helper.QuickSortRandom(v, 0, n-1)
	end := time.Since(start)

	fmt.Println(end)
}
