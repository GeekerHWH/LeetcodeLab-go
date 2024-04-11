package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	middle := (right-left)/2 + left //偶数个偏左

	for left <= right {
		// 左闭右开，相等返回
		if target == nums[middle] {
			return middle
		}
		if target > nums[middle] {
			left = middle + 1
		} else if target < nums[middle] {
			right = middle
		}

		middle = (right-left)/2 + left
		if left == right {
			if target > nums[middle] {
				return middle + 1
			} else if target < nums[middle] {
				return middle
			}
		}
	}
	return middle
}

func main() {
	// test case: nums, target := []int{1, 3, 5, 6}, 7 ->4
	// test case: nums, target := []int{1, 3, 5, 6}, 5 ->2
	// test case: nums, target := []int{1, 3, 5, 6}, 2 ->1
	// test case: nums, target := []int{1, 3, 5, 6}, 0 ->0

	nums, target := []int{1, 3, 5, 6}, 7
	pos := searchInsert(nums, target)
	fmt.Printf("%d", pos)
}
