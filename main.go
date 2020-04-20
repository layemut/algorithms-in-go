package main

import (
	"fmt"

	"github.com/layemut/algorithms-in-go/interview"
	"github.com/layemut/algorithms-in-go/search"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("index: ", search.Binary(arr, 4))
	fmt.Println("CountingValleys result: ", interview.CountingValleys(8, "UDDDUUUUDDUUUD"))
	fmt.Println("JumpingClouds result: ", interview.JumpingOnClouds([]int32{0, 1, 0, 1, 0, 1, 0}))
	fmt.Println("RepeatedString result: ", interview.RepeatedString("aaa", 21))
	fmt.Println("HourGlass result: ", interview.HourglassSum([][]int32{{0, -4, -6, 0, -7, -6}, {-1, -2, -6, -8, -3, -1}, {-8, -4, -2, -8, -8, -6}, {-3, -1, -2, -5, -7, -4}, {-3, -5, -3, -6, -6, -6}, {-3, -6, 0, -8, -6, -7}}))
	fmt.Println("RotateLeft result: ", interview.RotLeft([]int{1, 2, 3, 4}, 2))
}
