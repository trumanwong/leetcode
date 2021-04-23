package specialArray

import "sort"

func SpecialArray(nums []int) int {
	sort.Ints(nums)
	if nums[0] >= len(nums) {
		return len(nums)
	}

	for i := 1; i < len(nums); i++ {
		if nums[len(nums)-i] >= i && nums[len(nums)-i-1] < i {
			return i
		}
	}
	return -1
}
