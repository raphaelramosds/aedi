package helper

func LSearch(v [10]int, value int) int {
	for index, target := range v {
		if target == value {
			return index
		}
	}
	return -1
}
