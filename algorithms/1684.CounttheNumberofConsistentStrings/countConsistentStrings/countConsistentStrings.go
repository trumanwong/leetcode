package countConsistentStrings

func countConsistentStrings(allowed string, words []string) int {
	m := make(map[byte]bool)
	for i := 0; i < len(allowed); i++ {
		m[allowed[i]] = true
	}

	res := 0
	for _, word := range words {
		valid := true
		for i := 0; i < len(word); i++ {
			if _, ok := m[word[i]]; !ok {
				valid = false
				break
			}
		}
		if valid {
			res++
		}
	}
	return res
}