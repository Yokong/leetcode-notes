package main

import (
	"fmt"
	"leetcode-notes/src"
)

func main() {
	root := src.TreeNode{
		Val: 3,
	}
	root.Left = &src.TreeNode{
		Val: 9,
	}
	right := &src.TreeNode{
		Val: 20,
		Left: &src.TreeNode{
			Val: 15,
		},
		Right: &src.TreeNode{
			Val: 7,
		},
	}
	root.Right = right

	res := src.LevelOrderBottom(&root)
	fmt.Println(res)
}
