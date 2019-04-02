package main

import (
	"fmt"
	"leetcode-notes/src"
)

func main() {
	a := "(())"
	res := src.IsValid(a)
	fmt.Println(res)
}
