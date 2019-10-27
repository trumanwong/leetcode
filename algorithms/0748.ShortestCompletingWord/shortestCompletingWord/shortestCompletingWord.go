package shortestCompletingWord

func ShortestCompletingWord(licensePlate string, words []string) string {
	needle := make(map[rune]int)
	for _, v := range licensePlate {
		if v >= 'A' && v <='Z' {
			needle[v + 32]++
		} else if v >= 'a' && v <= 'z' {
			needle[v]++
		}
	}

	ret := ""
	for _, word := range words {
		m := make(map[rune]int)
		for _, v := range word {
			if _, ok := needle[v]; ok {
				m[v]++
			}
		}

		if len(m) < len(needle) {
			continue;
		}
		judge := true
		for k, v := range m {
			if needle[k] > v {
				judge = false
				break
			}
		}

		if !judge {
			continue
		}

		if ret == "" || len(ret) > len(word) {
			ret = word
		}
	}
	return ret
}