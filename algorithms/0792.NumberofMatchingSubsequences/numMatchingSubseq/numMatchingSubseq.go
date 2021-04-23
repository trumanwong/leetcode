package numMatchingSubseq

func NumMatchingSubseq(S string, words []string) int {
	count, sLen := 0, len(S)
	for _, word := range words {
		i, j, wLen := 0, 0, len(word)
		for i < wLen && j < sLen {
			if word[i] == S[j] {
				i++
				j++
			} else {
				j++
			}
		}
		if i == wLen {
			count++
		}
	}
	return count
}
