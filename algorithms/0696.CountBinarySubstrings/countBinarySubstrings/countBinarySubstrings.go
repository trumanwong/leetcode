package countBinarySubstrings

func CountBinarySubstrings(s string) int {
	last, curr, res := 0, 1, 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			curr++
		} else {
			last, curr = curr, 1
		}

		if last >= curr {
			res++
		}
	}
	return res
}
