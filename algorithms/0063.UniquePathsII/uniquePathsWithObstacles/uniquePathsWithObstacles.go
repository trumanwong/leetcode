package uniquePathsWithObstacles

func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid[0]), len(obstacleGrid)
	dp := make([]int, m + 1)
	dp[0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j] != 0 {
				dp[j] = 0
			} else if j != 0 {
				dp[j] += dp[j - 1]
			}
		}
	}
	return dp[len(dp) - 2]
}