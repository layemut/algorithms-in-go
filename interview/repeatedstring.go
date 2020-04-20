package interview

// Complete the repeatedString function below.
func RepeatedString(s string, n int64) int64 {
	acount := int64(0)
	count := n / int64(len(s))
	remain := n - (count * int64(len(s)))

	for _, char := range s {
		if char == 'a' {
			acount++
		}
	}
	acount = acount * count
	for _, char := range s[:remain] {
		if char == 'a' {
			acount++
		}
	}
	return acount
}
