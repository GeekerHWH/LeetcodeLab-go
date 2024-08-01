package quicksort

// QuickSort implements quick sort using slices, with greater RAM usage
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

// QuickSortII implements quick sort using double pointer, with better performance
func QuickSortII(nums []int, l int, r int) {
	if l >= r {
		return
	}
	i, j := l-1, r+1
	partition := nums[(l+r)/2]
	for i < j {
		for i++; nums[i] < partition; i++ {
		}
		for j--; nums[j] > partition; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	QuickSortII(nums, l, j)
	QuickSortII(nums, j+1, r)
}
