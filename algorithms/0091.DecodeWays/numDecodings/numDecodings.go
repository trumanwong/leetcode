package numDecodings

import "strconv"

func NumDecodings(s string) int {
	dp := make([]int, len(s)+1)

	if s == "" || s[0] == '0' {
		return 0
	}
	dp[0], dp[1] = 1, 1
	for i := 2; i < len(dp); i++ {
		val, _ := strconv.Atoi(s[i-2 : i])
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}

		if val <= 26 && val >= 10 {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}
