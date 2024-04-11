package main

import (
	"fmt"
	"slices"
)

func rotate(nums []int, k int) {
	k %= len(nums)
	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}

func main() {
	nums, k := []int{1, 2, 3, 4, 5, 6, 7}, 3
	rotate(nums, k)
	fmt.Println(nums)
}
