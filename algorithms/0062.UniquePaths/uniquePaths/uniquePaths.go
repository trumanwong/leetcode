package uniquePaths

func UniquePaths(m int, n int) int {
	var dp [][]int
	for i := 0; i < m; i++ {
		temp := make([]int, n)
		dp = append(dp, temp)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
			}
		}
	}
	return dp[m - 1][n - 1]
}