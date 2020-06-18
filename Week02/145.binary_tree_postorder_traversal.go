package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var r []int
	if root == nil {
		return r
	}
	r = append(r, postorderTraversal(root.Left))
	r = append(r, postorderTraversal(root.Right))
	r = append(r, root.Val)
	return r
}
