package main

import (
	"fmt"
)

func main() {
	input := [...]int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(input[:]))
}

func maxArea(height []int) int {
	var max, minHeight, area int

	for i, j := 0, len(height)-1; i < j; {
		if height[i] < height[j] {
			minHeight = height[i]
			i++
		} else {
			minHeight = height[j]
			j--
		}
		area = minHeight * (j - i + 1)
		max = Max(max, area)
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
