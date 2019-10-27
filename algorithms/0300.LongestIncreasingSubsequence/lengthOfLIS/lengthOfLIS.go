package lengthOfLIS

func LengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	ret, dp := 1, make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		maxVal := 0

		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				maxVal = max(maxVal, dp[j])
			}
		}
		dp[i] = maxVal + 1
		ret = max(ret, dp[i])
	}
	return ret
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}