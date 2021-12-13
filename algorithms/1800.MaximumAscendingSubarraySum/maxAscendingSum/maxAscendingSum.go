package maxAscendingSum

import "math"

func maxAscendingSum(nums []int) int {
	res, sum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			sum += nums[i]
		} else {
			res = int(math.Max(float64(res), float64(sum)))
			sum = nums[i]
		}
	}
	res = int(math.Max(float64(res), float64(sum)))
	return res
}