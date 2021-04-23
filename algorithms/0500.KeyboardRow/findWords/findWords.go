package findWords

func FindWords(words []string) []string {
	keyboards := [][]byte{[]byte("qwertyuiop"), []byte("asdfghjkl"), []byte("zxcvbnm")}
	res := []string{}
	for _, word := range words {
		m := make(map[int]int)
		for _, v := range []byte(word) {
			for i, keyboard := range keyboards {
				if inArray(v, keyboard) {
					m[i]++
					break
				}
			}
		}
		if len(m) == 1 {
			res = append(res, word)
		}
	}
	return res
}

func inArray(needle byte, stack []byte) bool {
	for _, v := range stack {
		if v == needle || v-32 == needle {
			return true
		}
	}
	return false
}
