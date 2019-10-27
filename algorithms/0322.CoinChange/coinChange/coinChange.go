package coinChange

func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount + 1)
	// 初始化dp数组， 最初所有的金额最优解为 -1 （不可达到）
	for i := 0; i <= amount; i++ {
		dp[i] = -1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i - coins[j] >= 0 && dp[i-coins[j]] != -1 && (dp[i] == -1 || dp[i] > dp[i - coins[j]] + 1) {
				dp[i] = dp[i-coins[j]] + 1;
			}
		}
	}
	return dp[amount]
}