package main

import (
	"fmt"
	"slices"
)

func threeSum(nums []int) [][]int {
	record := make(map[int]int, len(nums))

	// mapping (index and value) in nums as (value and key) in record
	for index, value := range nums {
		record[value] = index
	}

	// used for deduplicating
	intermediate := make(map[[3]int]bool)

	for i, val := range nums {
		for j, valu := range nums {
			if j == i { // i can't equal to j
				continue
			}

			if index, ok := record[0-val-valu]; ok && index != i && index != j {
				sortedArray := sort([]int{val, valu, 0 - val - valu})
				intermediate[sortedArray] = true
			}
		}
	}

	// deduplicating
	output := make([][]int, 0)
	for k, _ := range intermediate {

		output = append(output, k[:])
	}

	return output

}

func sort(num []int) [3]int {
	slices.Sort(num)
	array := [3]int{}
	copy(array[:], num)
	return array
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
