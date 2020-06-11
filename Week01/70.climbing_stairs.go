package main

import (
	"fmt"
	"math"
)

func main() {
	// input := 2
	// input := 4
	input := 450
	fmt.Println(climbStairs_1(input))
}

// Time:  O(2^n)
// Space: O(1)
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

// Time:  O(n)
// Space: O(1)
func climbStairs_1(n int) int {
	if n <= 2 {
		return n
	}
	c1, c2 := 1, 2
	for i := 3; i < n+1; i++ {
		c1, c2 = c2, c1+c2
	}
	return c2
}

// Time:  O(logn)
// Space: O(1)
func climbStairs_2(n int) int {
	sqrt_5 := math.Sqrt(5)
	fib_n := math.Pow((1+sqrt_5)/2, float64(n+1)) - math.Pow((1-sqrt_5)/2, float64(n+1))
	return int(fib_n / sqrt_5)
}
