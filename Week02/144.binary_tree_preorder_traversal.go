package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	r := append([]int{root.Val}, preorderTraversal(root.Left)...)
	r = append(r, preorderTraversal(root.Right)...)
	return r
}
