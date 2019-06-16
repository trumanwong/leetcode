package backspaceCompare

func BackspaceCompare(S string, T string) bool {
	s, t := []byte(S), []byte(T)
	for i := 1; i < len(s); i++ {
		if s[i] == '#' {
			key := i - 1
			for s[i] == '#' && key >= 0 {
				if s[key] != '#' {
					s[key] = '#'
					break
				}
				key--
			}
		}
	}
	for i := 1; i < len(t); i++ {
		key := i - 1
		for t[i] == '#' && key >= 0 {
			if t[key] != '#' {
				t[key] = '#'
				break
			}
			key--
		}
	}
	S, T = "", ""
	for _, v := range s {
		if v != '#' {
			S = S + string(v)
		}
	}
	for _, v := range t {
		if v != '#' {
			T = T + string(v)
		}
	}
	return S == T
}