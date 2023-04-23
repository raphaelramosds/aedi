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
	for i := 0; i < n; i++ {
		lowest_idx := 0
		for j := 0; j < n; j++ {
			if (*v)[lowest_idx] > (*v)[j] {
				lowest_idx = j
			}
		}
		sorted_list[i] = (*v)[lowest_idx]
		(*v)[lowest_idx] = max_int
	}
	*v = sorted_list
}

func SelectionSortIP(v []int) { // In place
	n := len(v)
	for i := 0; i < n-1; i++ {
		lowest_idx := i
		for j := i; j < n; j++ {
			if v[lowest_idx] > v[j] {
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
	for i := 0; i < n-1; i++ {
		swap := false
		for j := 0; j < n-i-1; j++ {
			if v[j] > v[j+1] {
				v[j], v[j+1] = v[j+1], v[j]
				swap = true
			}
		}
		if !swap {
			return
		}
	}
}

/*
* Insertion Sort
* Best case: Ω(n)
* Worst case: O(n²)
* Goal: insert the i-th element at its correct position on the "left hand"
 */

func InsertionSortIPV4(v []int) {
	n := len(v)
	for i := 1; i < n; i++ {
		for j := i; j > 0 && v[j] < v[j-1]; j-- {
			v[j], v[j-1] = v[j-1], v[j]
		}
	}
}

func InsertionSortIPV5(v []int) {
	n := len(v)
	for i := 1; i < n; i++ {
		var j int

		// Save the element to be inserted
		value := v[i]

		// Shift to right the elements of the left hand
		for j = i; j > 0 && value < v[j-1]; j-- {
			v[j] = v[j-1]
		}

		// Insert it at the correct position on left hand
		v[j] = value
	}
}
