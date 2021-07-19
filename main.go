package main

import (
	"fmt"
	"leetcode-notes/greedy"
)

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	res := greedy.CanCompleteCircuit(gas, cost)
	fmt.Println(res)
}

func AllForLoop() {
	a := []int{1, 2, 3}
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	fmt.Println("----------")
	for i := len(a) - 1; i >= 0; i-- {
		fmt.Println(a[i])
	}
	fmt.Println("----------")
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		fmt.Println(a[i], a[j])
	}
	fmt.Println("----------")
	for i := 1; i < len(a); i++ {
		fmt.Println(a[i])
	}
	fmt.Println("----------")
	for i := len(a) - 2; i >= 0; i-- {
		fmt.Println(a[i])
	}
}
