package interview

// Complete the countingValleys function below.
func CountingValleys(n int32, s string) int32 {
	valleycount := 0
	direction := 0
	runes := []rune(s)

	if runes[0] == 'U' {
		direction = direction + 1
	} else {
		direction = direction - 1
	}

	for i := 1; i < len(runes); i++ {
		if runes[i] == 'U' {
			direction = direction + 1
			if direction == 0 {
				valleycount++
			}
		} else {
			direction = direction - 1
		}
	}

	return int32(valleycount)
}
