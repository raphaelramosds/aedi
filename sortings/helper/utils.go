package helper

import (
	"fmt"
	"math/rand"
)

// Auxiliar functions

func Display(v []int) {
	fmt.Printf("[")
	for _, element := range v {
		fmt.Printf("%d ", element)
	}
	fmt.Printf("]\n")
}

func Init(length int) []int {
	list := make([]int, length)
	for i := 0; i < length; i++ {
		list[i] = rand.Intn(100)
	}
	return list
}

/*
* Selection Sort
* Best case: O(n)
* Worst case: O(n^2)
 */

func SelectionSort(v []int) {
	n := len(v)
	min_index := 0
	temp := 0
	for i := 0; i < n-1; i++ {
		for j := i; j < n; j++ {
			if v[temp] >= v[j] {
				temp = j
			}
		}
		v[min_index], v[temp] = v[temp], v[min_index]
		min_index++
	}
}
