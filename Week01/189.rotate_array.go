package main

import (
	"fmt"
)

func main() {
	input := [...]int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate_1(input[:], k)
	fmt.Println(input)
}

// Time:  O(1)
// Space: O(1)
func rotate(nums []int, k int) {
	k %= len(nums)
	copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))
}

// Time:  O(kn)
// Space: O(1)
func rotate_1(nums []int, k int) {
	for j := 1; j <= k; j++ {
		tmp := nums[len(nums)-1]
		for i, v := range nums {
			nums[i], tmp = tmp, v
		}
	}
}
