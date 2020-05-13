package main

import (
	"fmt"
	"leetcode-notes/src/array"
)

func main() {
	a := []int{1, 2, 3}
	array.NextPermutation(a)
	fmt.Println(a)
}
