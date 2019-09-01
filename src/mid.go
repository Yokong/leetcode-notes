package src

import (
	"leetcode-notes/stack"
)

// 98. 验证二叉搜索树
func IsValidBST(root *TreeNode) bool {
	stack := stack.New()
	inorder := -2147483648
	for stack.Len() > 0 || root != nil {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}

		root = stack.Pop().(*TreeNode)
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}

	return true
}
