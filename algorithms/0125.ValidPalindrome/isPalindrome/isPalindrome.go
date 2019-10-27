package isPalindrome

func IsPalindrome(s string) bool {
	i, j, sLen := 0, len(s) - 1, len(s)
	for i < j {
		for i < sLen && !((s[i] >= 48 && s[i] <= 57) || (s[i] >= 65 && s[i] <= 90) || (s[i] >= 97 && s[i] <= 122)) {
			i++
		}
		for j >= 0 && !((s[j] >= 48 && s[j] <= 57) || (s[j] >= 65 && s[j] <= 90) || (s[j] >= 97 && s[j] <= 122)) {
			j--
		}
		if i >= sLen || j < 0 {
			break
		}
		iTemp, jTemp := s[i], s[j]
		if iTemp >= 65 && iTemp <= 90 {
			iTemp +=32
		}
		if jTemp >= 65 && jTemp <= 90 {
			jTemp +=32
		}
		if iTemp != jTemp {
			return false
		}
		i++
		j--
	}
	return true
}