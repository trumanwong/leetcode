package arrayPairSum

import "sort"

func ArrayPairSum(nums []int) int {
	res := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i+= 2 {
		res += nums[i]
	}
	return res
}