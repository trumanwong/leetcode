package coinChange

import "math"

func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount + 1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] && dp[i - coins[j]] != math.MaxInt32 {
				dp[i] = int(math.Min(float64(dp[i - coins[j]] + 1), float64(dp[i])))
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}