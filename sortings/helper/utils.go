package helper

/*
* Selection Sort
* Best case: Ω(n²)
* Worst case: O(n²)
* Goal: move the lowest element to the left
 */

func SelectionSortOP(v *[]int) { // Out of place
	n := len(*v)
	max_int := 100
	sorted_list := make([]int, n)
	min_idx := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (*v)[min_idx] >= (*v)[j] {
				min_idx = j
			}
		}
		sorted_list[i] = (*v)[min_idx]
		(*v)[min_idx] = max_int
	}
	*v = sorted_list
}

func SelectionSortIP(v []int) { // In place
	n := len(v)
	for i := 0; i < n-1; i++ {
		lowest_idx := i
		for j := i; j < n; j++ {
			if v[lowest_idx] >= v[j] {
				lowest_idx = j
			}
		}
		v[lowest_idx], v[i] = v[i], v[lowest_idx]
	}
}

/*
* Bubble Sort
* Best case: Ω(n)
* Worst case: O(n²)
* Goal: move the highest element to the right
 */

func BubbleSort(v []int) {
	n := len(v)
	swap := false
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if v[j] >= v[j+1] {
				v[j], v[j+1] = v[j+1], v[j]
				swap = true
			}
		}
		if !swap {
			return
		}
	}
}
