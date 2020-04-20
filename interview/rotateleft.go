package interview

// Complete the rotLeft function below.
func RotLeft(a []int, d int) []int {
	left := a[:d]
	right := a[d:]
	return append(right, left...)
}
