package reverseOnlyLetters

func ReverseOnlyLetters(S string) string {
	l, r, res := 0, len(S) - 1, make([]byte, len(S))
	for l <= r {
		if !((S[l] >= 65 && S[l] <= 90) || (S[l] >= 97 && S[l] <= 122)) {
			res[l] = S[l]
			l++
		} else if !((S[r] >= 65 && S[r] <= 90) || (S[r] >= 97 && S[r] <= 122)) {
			res[r] = S[r]
			r--
		} else {
			res[l], res[r] = S[r], S[l]
			l++
			r--
		}
	}
	return string(res)
}