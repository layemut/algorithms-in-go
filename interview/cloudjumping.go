package interview

// Complete the jumpingOnClouds function below.
func JumpingOnClouds(c []int32) int32 {
	steps := 0
	clouds := len(c)
	for i := 0; i < clouds; i++ {
		if i+2 < clouds && c[i+2] != 1 {
			i = i + 1
			steps++
		} else if i+1 < clouds && c[i+1] != 1 {
			steps++
		}
	}
	return int32(steps)
}
