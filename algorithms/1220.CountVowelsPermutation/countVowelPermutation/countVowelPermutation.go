package countVowelPermutation

func CountVowelPermutation(n int) int {
	dp := [5]int{1, 1, 1, 1, 1}
	nx := [5]int{}
	mod := 1000000007
	for i := 2; i <= n; i++ {
		nx[0] = (dp[1] + dp[2] + dp[4]) % mod
		nx[1] = (dp[0] + dp[2]) % mod
		nx[2] = (dp[1] + dp[3]) % mod
		nx[3] = dp[2]
		nx[4] = (dp[2] + dp[3]) % mod
		dp, nx = nx, dp
	}
	return (dp[0] + dp[1] + dp[2] + dp[3] + dp[4]) % mod
}
