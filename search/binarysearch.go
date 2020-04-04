package search

// Binary binary serach
func Binary(arr []int, x int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		m := left + (right-left)/2

		if arr[m] == x {
			return m
		}

		if arr[m] < x {
			left = m + 1
		} else {
			right = m - 1
		}
	}
	return -1
}
