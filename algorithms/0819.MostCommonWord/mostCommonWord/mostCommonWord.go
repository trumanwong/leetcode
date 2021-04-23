package mostCommonWord

func MostCommonWord(paragraph string, banned []string) string {
	m, p := make(map[string]int), []byte(paragraph)
	start := -1
	for i, _ := range p {
		if p[i] >= 'A' && p[i] <= 'Z' {
			p[i] = p[i] + 32
		}

		if !(p[i] >= 'a' && p[i] <= 'z') && start != -1 {
			temp := string(p[start:i])
			m[temp]++
			start = -1
		} else if p[i] >= 'a' && p[i] <= 'z' && start == -1 {
			start = i
		}
	}
	if start != -1 {
		temp := string(p[start:len(p)])
		m[temp]++
	}
	max, res := 0, ""
	for k, v := range m {
		if v > max && !inArray(k, banned) {
			max = v
			res = k
		}
	}
	return res
}

func inArray(needle string, hayStack []string) bool {
	for _, v := range hayStack {
		if v == needle {
			return true
		}
	}
	return false
}
