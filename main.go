package main

import (
	"fmt"
	"leetcode-notes/stack"
)

func main() {
	s := stack.New()
	s.Push(10)
	s.Push(11)
	s.Push(12)
	s.Push(13)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
