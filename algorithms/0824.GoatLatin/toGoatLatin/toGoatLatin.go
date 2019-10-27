package toGoatLatin

func ToGoatLatin(S string) string {
	s, res := []byte(S), []byte{}
	start, count := -1, 0
	vowel := []byte("aeiouAEIOU")
	for i, v := range s {
		if v != ' ' && start == -1 {
			start = i
			count++
		}
		if v == ' ' && start != -1 {
			temp := []byte{}
			if !inArray(s[start], vowel) {
				temp = append(temp, s[start + 1:i]...)
				temp = append(temp, s[start])
			} else {
				temp = append(temp, s[start:i]...)
			}
			temp = append(temp, []byte("ma")...)
			for i := 0; i < count; i++ {
				temp = append(temp, 'a')
			}
			temp = append(temp, ' ')
			res = append(res, temp...)
			start = -1
		}
	}

	if start != -1 {
		temp := []byte{}
		if !inArray(s[start], vowel) {
			temp = append(temp, s[start + 1:len(s)]...)
			temp = append(temp, s[start])
		} else {
			temp = append(temp, s[start:len(s)]...)
		}
		temp = append(temp, []byte("ma")...)
		for i := 0; i < count; i++ {
			temp = append(temp, 'a')
		}
		temp = append(temp)
		res = append(res, temp...)
	}
	return string(res)
}

func inArray(needle byte, hayStack []byte) bool {
	for _, v := range hayStack {
		if v == needle {
			return true
		}
	}
	return false
}