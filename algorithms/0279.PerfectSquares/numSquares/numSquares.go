package numSquares

import "math"

func NumSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		temp := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			temp = min(temp, dp[i-j*j])
		}
		dp[i] = temp + 1
	}
	return dp[n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
