package longestArithSeqLength

func LongestArithSeqLength(A []int) int {
	m1, m2, res := make(map[int]int), make(map[int]int), 0
	for i, v := range A {
		for j := i + 1; j < len(A); j++ {
			diff := v - A[j]
			m1[diff]++
			if m1[diff] == 1 {
				m2[diff] = v
			}
			if m1[diff] > res && (m1[diff] == 1 || (diff < 0 && m2[diff] < A[j]) || (diff > 0 && m2[diff] > A[j])) {
				res = m1[diff]
			}
		}
	}
	return res
}