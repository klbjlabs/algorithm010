package main

import (
	"fmt"
)

func main() {
	// input := [...]int{1, 2, 3, 4}
	input := [...]int{9, 9, 9, 9}

	fmt.Println(plusOne(input[:]))
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}
	}
	nums := make([]int, len(digits)+1)
	nums[0] = 1
	return nums
}
