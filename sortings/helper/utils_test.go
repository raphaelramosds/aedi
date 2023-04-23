package helper

import (
	"math/rand"
	"testing"
)

// # of elements

var size int = 90000

/*
* @brief Inits an array of length n
* @param n int
 */

func create_list(n int) []int {
	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] = rand.Intn(100)
	}
	return list
}

/*
* @brief Verify if an array is ascendent sorted
* @param v []int
 */

func is_asc_sorted(v []int) bool {
	n := len(v)
	for i := 0; i < n-1; i++ {
		if v[i] > v[i+1] {
			return false
		}
	}
	return true
}

func TestSelectionSortIP(t *testing.T) {
	s := create_list(size)
	SelectionSortIP(s[:])
	result := is_asc_sorted(s[:])
	if !result {
		t.Errorf("SelectionSortIP did not sort properly")
	}
}

func TestSelectionSortOP(t *testing.T) {
	s := create_list(size)
	SelectionSortOP(&s)
	result := is_asc_sorted(s[:])
	if !result {
		t.Errorf("SelectionSortOP did not sort properly")
	}
}

func TestBubbleSort(t *testing.T) {
	s := create_list(size)
	BubbleSort(s[:])
	result := is_asc_sorted(s[:])
	if !result {
		t.Errorf("BubbleSort did not sort properly")
	}
}

func TestInsertionSortIPV4(t *testing.T) {
	s := create_list(size)
	InsertionSortIPV4(s[:])
	result := is_asc_sorted(s[:])
	if !result {
		t.Errorf("InsertionSortIPV4 did not sort properly")
	}
}

func TestInsertionSortIPV5(t *testing.T) {
	s := create_list(size)
	InsertionSortIPV5(s[:])
	result := is_asc_sorted(s[:])
	if !result {
		t.Errorf("InsertionSortIPV5 did not sort properly")
	}
}