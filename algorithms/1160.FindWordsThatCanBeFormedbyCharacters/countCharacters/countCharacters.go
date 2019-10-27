package countCharacters

func CountCharacters(words []string, chars string) int {
	m := make(map[rune]int)
	for _, c := range chars {
		m[c]++
	}
	ret := 0
	for _, word := range words {
		isContinue := false
		temp := make(map[rune]int)
		for _, v := range word {
			temp[v]++
			if _, ok := m[v]; !ok || temp[v] > m[v] {
				isContinue = true
				break
			}
		}

		if isContinue {
			continue
		}
		ret += len(word)
	}
	return ret
}