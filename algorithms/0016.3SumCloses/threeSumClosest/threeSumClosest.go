package threeSumClosest

import "sort"

func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		start, end := i + 1, len(nums) - 1
		for start < end {
			sum := nums[start] + nums[end] + nums[i]
			if abs(target - sum) < abs(target - res) {
				res = sum
			} else if sum > target {
				end--
			} else if sum < target {
				start++
			} else {
				return res
			}
		}
	}
	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}