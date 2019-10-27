package probabilityOfHeads

// 有一些不规则的硬币。在这些硬币中，prob[i] 表示第 i 枚硬币正面朝上的概率。
//请对每一枚硬币抛掷 一次，然后返回正面朝上的硬币数等于 target 的概率。
func ProbabilityOfHeads(prob []float64, target int) float64 {
	var dp [1005][1005]float64
	dp[0][0] = 1
	for i := 1; i <= len(prob); i++ {
		for j := 0; j <= len(prob); j++ {
			if j != 0 {
				dp[i][j] += dp[i-1][j-1] * prob[i-1]
			}
			dp[i][j] += dp[i-1][j] * (1-prob[i-1])
		}
	}
	return dp[len(prob)][target];
}