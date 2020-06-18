package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list := [...]int{1, 2, 3, 4, 5}
}

func initLinkedList(list []int) ListNode {
	var linkedList ListNode
	for _, v := range list {
		node := v
		linkedList.Next = node
		linkedList = 
	}
}
