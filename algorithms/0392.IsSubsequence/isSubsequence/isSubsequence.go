package isSubsequence

func IsSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}

	i, j, sLen, tLen := 0, 0, len(s), len(t)
	for i < sLen && j < tLen {
		if s[i] == t[j] {
			i++
			j++
		} else {
			j++
		}
	}
	if i == sLen {
		return true
	}
	return false
}