package uniqueMorseRepresentations

func UniqueMorseRepresentations(words []string) int {
	morse := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	m := make(map[string]int)
	for _, word := range words {
		temp := ""
		for _, v := range word {
			temp = temp + morse[v-'a']
		}
		m[temp]++
	}
	return len(m)
}
