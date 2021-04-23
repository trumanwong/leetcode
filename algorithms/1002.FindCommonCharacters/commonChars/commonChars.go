package commonChars

func CommonChars(A []string) []string {
	m, res := make(map[byte]int), []string{}
	for _, v := range []byte(A[0]) {
		m[v]++
	}
	for i := 1; i < len(A); i++ {
		temp := make(map[byte]int)
		for _, v := range []byte(A[i]) {
			temp[v]++
		}
		for i := range m {
			if _, ok := temp[i]; !ok || m[i] != temp[i] {
				if ok {
					m[i] = min(m[i], temp[i])
				} else {
					delete(m, i)
				}
			}
		}
	}
	for i, v := range m {
		for j := 0; j < v; j++ {
			res = append(res, string(i))
		}
	}
	return res
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
