package binaryGap

func BinaryGap(N int) int {
	last, ans := -1, 0
	for i := 0; i < 32; i++ {
		if (N>>uint(i))&1 > 0 {
			if last >= 0 {
				ans = max(ans, i-last)
			}
			last = i
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
