package quicksort

import "testing"

func TestQuickSort(t *testing.T) {
	nums := []int{2, 5, 8, 6, 9, 4, 10, 1}

	expected := []int{1, 2, 4, 5, 6, 8, 9, 10}
	result := QuickSort(nums)

	if !sliceEqual(result, expected) {
		t.Errorf("QuickSort failed, got %v, expected %v", nums, expected)
	}
}

func TestQuickSortII(t *testing.T) {
	nums := []int{2, 5, 8, 6, 9, 4, 10, 1}

	expected := []int{1, 2, 4, 5, 6, 8, 9, 10}
	QuickSortII(nums, 0, len(nums)-1)

	if !sliceEqual(nums, expected) {
		t.Errorf("QuickSort failed, got %v, expected %v", nums, expected)
	}
}

// Used for slices comparison
func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
