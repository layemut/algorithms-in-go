package main

import (
	"github.com/layemut/algorithms-in-go/search"
	"fmt"
)
func main(){
  arr := []int{1,2,3,4,5}
  fmt.Println("index: ", search.Binary(arr, 2))
}