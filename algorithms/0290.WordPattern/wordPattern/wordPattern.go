package wordPattern

func WordPattern(pattern string, str string) bool {
	start ,strs := -1, []string{}
	p, b := []byte(pattern), []byte(str)
	for i, v := range b {
		if v == ' ' && start != -1 {
			strs = append(strs, string(b[start:i]))
			start = -1
		} else if start == -1 {
			start = i
		}
	}
	if start != -1 {
		strs = append(strs, string(b[start:len(b)]))
	}
	if len(strs) != len(pattern) {
		return false
	}
	m1, m2 := make(map[byte]string), make(map[string]byte)
	for i, v := range p {
		if _, ok := m1[v]; !ok {
			m1[v] = strs[i]
		} else {
			if m1[v] != strs[i] {
				return false
			}
		}
		if _, ok := m2[strs[i]]; !ok {
			m2[strs[i]] = v
		} else {
			if m2[strs[i]] != v {
				return false
			}
		}
	}

	return true
}