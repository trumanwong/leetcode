package frequencySort

func FrequencySort(s string) string {
	m := make(map[rune]int)

	S := []rune{}
	for _, v := range s {
		if _, ok := m[v]; !ok {
			S = append(S, v)
		}
		m[v]++
	}

	for i := 0; i < len(S); i++ {
		for j := i + 1; j < len(S); j++ {
			if m[S[i]] < m[S[j]] {
				S[i], S[j] = S[j], S[i]
			}
		}
	}

	ret := []rune{}
	for _, v := range S {
		for i := 0; i < m[v]; i++ {
			ret = append(ret, v)
		}
	}
	return string(ret)
}
