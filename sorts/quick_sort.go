package quicksort

func QuickSort(nums []int) []int {
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

	less = QuickSort(less)
	greater = QuickSort(greater)

	return append(append(less, equal...), greater...)
}
