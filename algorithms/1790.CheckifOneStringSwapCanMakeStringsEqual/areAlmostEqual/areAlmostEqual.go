package areAlmostEqual

func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	index, diff := -1, 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff++
			if diff == 1 {
				index = i
			} else if diff > 2 || s1[index] != s2[i] || s2[index] != s1[i] {
				return false
			}
		}
	}
	return diff == 2
}
