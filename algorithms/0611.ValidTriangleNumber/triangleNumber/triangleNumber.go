package triangleNumber

import "sort"

func TriangleNumber(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i := len(nums) - 1; i > 1; i-- {
		l, r := 0, i - 1
		for l < r {
			// 只要较小的两个值之和大于最大值, 则一定可以组成三角形
			if nums[l] + nums[r] > nums[i] {
				res += r - l
				r--
			} else {
				l++
			}
		}
	}
	return res
}