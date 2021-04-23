package findUnsortedSubarray

import "sort"

func FindUnsortedSubarray(nums []int) int {
	nums2 := make([]int, len(nums))
	copy(nums2, nums)
	sort.Ints(nums2)

	start, end := len(nums), 0
	for i, v := range nums2 {
		if v != nums[i] {
			start = min(start, i)
			end = max(end, i)
		}
	}

	if end-start >= 0 {
		return end - start + 1
	}
	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
