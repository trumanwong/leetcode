package isMatch

func IsMatch(s string, p string) bool {
	if p == "" {
		return s == p
	}

	firstMatch := false
	if s != "" && p[0] == s[0] || p[0] == '.' {
		firstMatch = true
	}

	if len(p) >= 2 && p[1] == '*' {
		return IsMatch(s, p[2:]) || (firstMatch && len(s) > 0 && IsMatch(s[1:], p))
	}
	return firstMatch && len(s) > 0 && len(p) > 0 && IsMatch(s[1:], p[1:])
}