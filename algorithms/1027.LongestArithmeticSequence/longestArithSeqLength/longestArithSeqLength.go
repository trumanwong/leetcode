package longestArithSeqLength

func LongestArithSeqLength(A []int) int {
	if len(A) <= 2 {
		return len(A)
	}
	min, max := 10001, 0
	for _, v := range A {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}
	diff, res := max - min, 1
	//dp[i][j]表示前i个元素 子序列差值为j的子序列的长度
	dp := make([][]int, len(A))

	for i, v := range A {
		dp[i] = make([]int, 2 * diff + 1)
		for j := 0; j < i; j++ {
			temp := v - A[j] + diff
			dp[i][temp] = dp[j][temp] + 1
			res = Max(res, dp[i][temp])
		}
	}

	return res + 1
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}