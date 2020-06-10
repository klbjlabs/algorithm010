package main

import (
	"fmt"
)

func main() {
	input := [...]int{0, 1, 0, 3, 12}
	moveZeros(input[:])
	fmt.Println(input)
}

// TimeComplexity:  O(n)
// SpaceComplexity: O(1)
func moveZeros(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
