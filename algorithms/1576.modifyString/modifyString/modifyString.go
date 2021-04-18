package modifyString

func ModifyString(s string) string {
	b := []rune(s)
	if len(b) == 1 && b[0] == '?' {
		return "a"
	}
	for i, v := range b {
		if v != '?' {
			continue
		}
		if i == 0 {
			b[i] = validCharacter(b[i+1], b[i+1])
		} else if i == len(b)-1 {
			b[i] = validCharacter(b[i-1], b[i-1])
		} else {
			b[i] = validCharacter(b[i-1], b[i+1])
		}
	}
	return string(b)
}

func validCharacter(l, r rune) rune {
	res := 'a'
	for i := 'a'; i <= 'z'; i++ {
		if i != l && i != r {
			res = i
			break
		}
	}
	return res
}
