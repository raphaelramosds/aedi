package helper

func BSearch(start int, end int, v [5]int, value int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	if v[mid] == value {
		return mid
	} else {
		if value < v[mid] {
			return BSearch(start, mid-1, v, value)
		} else {
			return BSearch(mid+1, end, v, value)
		}
	}
}
