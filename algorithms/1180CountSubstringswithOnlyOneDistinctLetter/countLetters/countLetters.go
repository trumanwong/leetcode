package countLetters

func countLetters(S string) int {
	last := S[0]
	ret, length := 1, 1
	for i := 1; i < len(S); i++ {
		if S[i] != last {
			length = 1
			last = S[i]
		} else {
			length++
		}
		ret += length
	}
	return ret
}