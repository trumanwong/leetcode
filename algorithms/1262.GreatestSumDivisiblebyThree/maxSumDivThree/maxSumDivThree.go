package maxSumDivThree

func MaxSumDivThree(nums []int) int {
	// 设dp[i] 代表 选取的数字累加和 模3 = i 的数字和
	// 假定nums[i] % 3 = 1 ，那么，和 前面选取的数字和模 3 = 2 的数相加，就可以模3为 0 ，表达起来就是 dp[0] = max(dp[0], nums[i] + dp[2])
	// 依次类推，只要不断更新 dp 数组即可，注意一点，更新的时候要保存上一个状态的值，避免后续更新的时候重复影响。
	dp := []int{0, 0, 0}
	for _, v := range nums {
		mod := v % 3

		a := dp[(3+0-mod)%3]
		b := dp[(3+1-mod)%3]
		c := dp[(3+2-mod)%3]

		if a != 0 || mod == 0 {
			dp[0] = max(dp[0], a+v)
		}
		if b != 0 || mod == 1 {
			dp[1] = max(dp[1], b+v)
		}
		if c != 0 || mod == 2 {
			dp[2] = max(dp[2], c+v)
		}
	}
	return dp[0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
