package maxSumAfterPartitioning

func MaxSumAfterPartitioning(A []int, K int) int {
	res := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		end := i + K - 1
		if end >= len(A) {
			end = len(A) - 1
		}

		for j := i; j <= end; j++ {
			start := j - K + 1
			if start < 0 {
				start = 0
			}

			for t := start; t <= i; t++ {
				pre := 0
				if t - 1 >= 0 {
					pre = res[t - 1]
				}
				res[j] = max(res[j], pre + A[i] * (j - t + 1))
			}
		}
	}
	return res[len(A) - 1]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}