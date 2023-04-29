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

/*
* Merge Sort
* Best case: Ω(nlog(n))
* Worst case: O(nlog(n))
* Goal: divide and conquer
 */

func Merge(v, v1, v2 []int) {
	i1 := 0
	i2 := 0
	k := 0
	l1 := len(v1)
	l2 := len(v2)

	for i1 < l1 && i2 < l2 {
		if v1[i1] > v2[i2] {
			v[k] = v2[i2]
			i2++
		} else {
			v[k] = v1[i1]
			i1++
		}
		k++
	}

	for i1 < l1 {
		v[k] = v1[i1]
		i1++
		k++
	}

	for i2 < l2 {
		v[k] = v2[i2]
		i2++
		k++
	}
}

func MergeSort(v []int) {
	if len(v) > 1 {

		// Find the mid index

		mid := len(v) / 2

		// Split the left and right halves of the array

		left_half := make([]int, mid)
		right_half := make([]int, len(v)-mid)

		// Populate halves

		for lh := 0; lh < mid; lh++ {
			left_half[lh] = v[lh]
		}

		for rh := mid; rh < len(v); rh++ {
			right_half[rh-mid] = v[rh]
		}

		// Recursion calls

		MergeSort(left_half[:])
		MergeSort(right_half[:])

		// Merge two halves

		Merge(v[:], left_half[:], right_half[:])
	}
}
