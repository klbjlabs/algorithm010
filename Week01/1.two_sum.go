package main

import (
	"fmt"
)

func main() {
	xx := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}
	// ages := make(map[string]int)
	// delete(ages, "alice")

	for k, v := range xx {
		fmt.Println(k, v)
	}
}
