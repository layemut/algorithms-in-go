package interview

// Complete the hourglassSum function below.
func HourglassSum(arr [][]int32) int32 {
	sum := int32(0)
	imid := int32(len(arr) / 2)
	jmid := int32(len(arr[0]) / 2)
	max := arr[imid][jmid-1] + arr[imid][jmid] + arr[imid][jmid+1] + arr[imid+1][jmid] + arr[imid+2][jmid-1] + arr[imid+2][jmid] + arr[imid+2][jmid+1]
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			sum = 0
			if j-1 >= 0 && j+1 < len(arr[i]) && i+2 < len(arr) {
				sum = sum + arr[i][j-1] + arr[i][j] + arr[i][j+1]
				sum = sum + arr[i+1][j]
				sum = sum + arr[i+2][j-1] + arr[i+2][j] + arr[i+2][j+1]
				if sum > max {
					max = sum
				}
			}
		}
	}
	return max
}
