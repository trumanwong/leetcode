package topKFrequent

func TopKFrequent(words []string, k int) []string {
	m := make(map[string]int)

	ws := []string{}
	for _, word := range words {
		if _, ok := m[word]; !ok {
			ws = append(ws, word)
		}
		m[word]++
	}

	ret := []string{}
	for i := 0; i < len(ws); i++ {
		for j := i + 1; j < len(ws); j++ {
			if m[ws[i]] < m[ws[j]] {
				ws[i], ws[j] = ws[j], ws[i]
			} else if m[ws[i]] == m[ws[j]] {
				is_equal := false
				for t := 0; t < len(ws[i]) && t < len(ws[j]); t++ {
					if ws[i][t] > ws[j][t] {
						is_equal = false
						ws[i], ws[j] = ws[j], ws[i]
						break
					} else if ws[i][t] < ws[j][t] {
						is_equal = false
						break
					} else {
						is_equal = true
					}
				}

				if is_equal && len(ws[i]) > len(ws[j]) {
					ws[i], ws[j] = ws[j], ws[i]
				}
			}
		}

		ret = append(ret, ws[i])
		if len(ret) == k {
			break
		}
	}

	return ret
}