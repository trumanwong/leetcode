package minSwaps

func MinSwaps(data []int) int {
	n := len(data)
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		s[i] = s[i-1] + data[i-1]
	}
	m, ret := s[n], n
	for i := m; i <= n; i++ {
		ret = min(ret, m-(s[i]-s[i-m]))
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
