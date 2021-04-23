package reverseWords

func ReverseWords(s string) string {
	res, start := []byte(s), 0
	for i, v := range s {
		if v == ' ' && i > start {
			reverse(res[start:i])
			start = i + 1
		}
	}
	reverse(res[start:])
	return string(res)
}

func reverse(s []byte) {
	i, j := 0, len(s)-1
	for i <= j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
