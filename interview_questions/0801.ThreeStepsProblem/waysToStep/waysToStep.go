package waysToStep

func WaysToStep(n int) int {
	if n < 3 {
		return n
	}
	if n == 3 {
		return 4
	}
	dp := make([]int, n)
	dp[0], dp[1], dp[2] = 1, 2, 4
	for i := 3; i < n; i++ {
		dp[i] = (dp[i - 1] + dp[i - 2]) % 1000000007 + dp[i - 3]
		dp[i] %= 1000000007
	}
	return dp[n - 1]
}