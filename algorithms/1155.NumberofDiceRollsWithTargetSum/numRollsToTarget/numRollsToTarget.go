package numRollsToTarget

func NumRollsToTarget(d int, f int, target int) int {
	var dp [][]int
	for i := 0; i < 31; i++ {
		temp := make([]int, 1001)
		dp = append(dp, temp)
	}
	dp[0][0] = 1
	for i := 1; i <= d; i++ {
		for k := 1; k <= f; k++ {
			temp := f * d
			for j := 1; j <= temp; j++ {
				if j >= k {
					dp[i][j] += dp[i-1][j-k]
					dp[i][j] %= 1000000007
				}
			}
		}
	}

	return dp[d][target]
}
