package maxRepOpt1

func MaxRepOpt1(text string) int {
	s, m := []byte(text), make(map[byte]int)
	for _, v := range s {
		m[v]++
	}

	contious, ret := 1, 0
	curr, flag := s[0], false
	i, diff := 1, -1
	for i < len(s) {
		if s[i] == curr {
			contious++
			if m[curr] < contious {
				contious = m[curr]
			}
		} else {
			if !flag {
				diff = i
				flag = true
				contious++
				if m[curr] < contious {
					contious = m[curr]
				}
			} else {
				i, flag = diff, false
				if contious > ret {
					ret = contious
				}
				contious = 1
				curr = s[i]
			}
		}
		i++
	}

	if contious > ret {
		ret = contious
	}
	return ret
}