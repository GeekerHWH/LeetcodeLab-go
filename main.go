package main

import (
	"fmt"
	"math"
)

func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	pivot := nums[0]
	less := []int{}
	equal := []int{pivot}
	greater := []int{}

	for _, num := range nums[1:] {
		if num < pivot {
			less = append(less, num)
		} else if num > pivot {
			greater = append(greater, num)
		} else {
			equal = append(equal, num)
		}
	}

	less = quickSort(less)
	greater = quickSort(greater)

	return append(append(less, equal...), greater...)
}

func minSubArrayLen(target int, nums []int) int {
	total := sum(nums)
	if total < target {
		return 0
	} else if total == target {
		return len(nums)
	}

	result := math.MaxInt32
	slow, fast := 0, 0
	for fast < len(nums) && slow <= fast {
		sumOfWindow := sum(nums[slow : fast+1])
		if sumOfWindow < target {
			fast++
			continue
		}
		result = min(result, fast+1-slow)
		slow++
	}
	return result
}

func sum(nums []int) int {
	result := 0
	for _, val := range nums {
		result += val
	}
	return result
}

func min(a int, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}

func main() {
	nums := []int{8}
	target := 7
	fmt.Println(minSubArrayLen(target, nums))
}
