package lastSubstring

func LastSubstring(s string) string {
	max := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
		}
	}

	ret := ""
	for i := 0; i < len(s); i++ {
		if s[i] == max {
			ret = Max(ret, s[i:])
		}
	}
	return ret
}

func Max(a, b string) string {
	if a > b {
		return a
	}
	return b
}