package helper

import "math/rand"

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
		for j := i + 1; j < n; j++ {
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
* Goal: split array into halves with recursion calls. Then,
* given two halves that are sorted, merge them into a sorted array
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

/*
* Quick Sort/Partition Sort
* Best case: Ω(nlog(n))
* Worst case: O(n²) if pivot isn't randomly taken, otherwise, O(nlog(n))
* Goal: on each iteration, choose a pivot and put elements greater than it
* at its right and lower ones at its left
 */

func QuickSort(v []int, left int, right int) {
	if left < right {
		pivot_index := PartitionV1(v, left, right)
		QuickSort(v, left, pivot_index-1)
		QuickSort(v, pivot_index+1, right)
	}
}

func QuickSortRandomPivot(v []int, left int, right int) {
	if left < right {
		pivot_index := PartitionV2(v, left, right)
		QuickSort(v, left, pivot_index-1)
		QuickSort(v, pivot_index+1, right)
	}
}

// Without randomization: pivot index is at the end of the partition
func PartitionV1(v []int, left int, right int) int {
	new_pindex := left
	pivot := v[right]
	for i := left; i < right; i++ {
		if v[i] < pivot {
			v[i], v[new_pindex] = v[new_pindex], v[i]
			new_pindex++
		}
	}
	v[new_pindex], v[right] = v[right], v[new_pindex]
	return new_pindex
}

// With randomization: pivot index is randomly choosen
func PartitionV2(v []int, left, right int) int {

	// Generate a random index between left and right
	random_index := rand.Intn(right-left) + left

	// Put the element v[random_index] at the end
	v[random_index], v[right] = v[right], v[random_index]

	// Proceeds naturally
	new_pindex := left
	pivot := v[right]
	for i := left; i < right; i++ {
		if v[i] < pivot {
			v[i], v[new_pindex] = v[new_pindex], v[i]
			new_pindex++
		}
	}
	v[new_pindex], v[right] = v[right], v[new_pindex]
	return new_pindex
}

/*
* Counting Sort
* Best case: Ω(n)
* Worst case: O(n+k), k >> n
* Goal: sorting integer arrays by counting
**/

func CountingSort(v []int) []int {

	// Find the lowest and highest element
	lowest := v[0]
	highest := v[0]
	for i := 1; i < len(v); i++ {
		if v[i] > highest {
			highest = v[i]
		}
		if v[i] < lowest {
			lowest = v[i]
		}
	}

	// Create counting array
	c := make([]int, highest-lowest+1)

	// Populate counting array with the frequency of each element
	for i := 0; i < len(v); i++ {
		c[v[i]-lowest]++
	}

	// Cumulative sum over counting array
	for i := 1; i < len(c); i++ {
		c[i] += c[i-1]
	}

	// Create sorted array
	s := make([]int, len(v))

	// Find the equivalent position of v[i] in array c and insert it on array s
	for i := 0; i < len(v); i++ {
		s[c[v[i]-lowest]-1] = v[i]
		c[v[i]-lowest]--
	}

	// Return the sorted array
	return s
}
