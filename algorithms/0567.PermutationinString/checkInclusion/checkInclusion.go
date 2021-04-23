package checkInclusion

func CheckInclusion(s1 string, s2 string) bool {
	cnt, start := 0, 0
	m := make(map[byte]int)
	for _, v := range []byte(s1) {
		m[v]++
	}

	for i := 0; i < len(s2) && cnt < len(s1); i++ {
		if m[s2[i]] > 0 {
			m[s2[i]]--
			cnt++
		} else {
			for s2[start] != s2[i] {
				if _, ok := m[s2[start]]; ok {
					m[s2[start]]++
					cnt--
				}
				start++
			}
			start++
		}
	}
	return cnt == len(s1)
}
