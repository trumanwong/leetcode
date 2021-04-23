package largestDivisibleSubset

import (
	"math"
	"sort"
)

func LargestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)

	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	maxSize, maxVal := 1, dp[0]
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
			}
		}

		if dp[i] > maxSize {
			maxSize, maxVal = dp[i], nums[i]
		}
	}

	// 倒推获得最大子集
	res := make([]int, 0)
	if maxSize == 1 {
		res = append(res, nums[0])
		return res
	}

	for i := len(nums) - 1; i >= 0 && maxSize > 0; i-- {
		if dp[i] == maxSize && maxVal%nums[i] == 0 {
			res = append(res, nums[i])
			maxVal = nums[i]
			maxSize--
		}
	}
	return res
}
