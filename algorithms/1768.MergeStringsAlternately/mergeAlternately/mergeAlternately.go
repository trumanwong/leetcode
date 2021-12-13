package mergeAlternately

func mergeAlternately(word1 string, word2 string) string {
	res := make([]byte, len(word1) + len(word2))
	i, j := 0, 0
	for i < len(word1) || i < len(word2) {
		if i < len(word1) {
			res[j] = word1[i]
			j++
		}
		if i < len(word2) {
			res[j] = word2[i]
			j++
		}
		i++
	}
	return string(res)
}