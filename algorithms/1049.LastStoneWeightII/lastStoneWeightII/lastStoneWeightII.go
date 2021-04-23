package lastStoneWeightII

func LastStoneWeightII(stones []int) int {
	sum := 0
	for _, v := range stones {
		sum += v
	}

	dp := make([]int, 1501)
	for _, v := range stones {
		for j := sum / 2; j >= v; j-- {
			dp[j] = max(dp[j], dp[j-v]+v)
		}
	}
	return sum - 2*dp[sum/2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
