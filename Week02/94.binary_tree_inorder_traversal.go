package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	r := []int{}

	if root == nil {
		return r
	}
	r = append(r, inorderTraversal(root.Left)...)
	r = append(r, root.Val)
	r = append(r, inorderTraversal(root.Right)...)
	return r
}
