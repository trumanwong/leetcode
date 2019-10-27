package removeDuplicateLetters

func RemoveDuplicateLetters(s string) string {
	count, show := make([]int, 26), make([]int, 26)
	for _, v := range s {
		count[v - 'a']++
	}

	ret := make([]rune, 0)
	for _, v := range s {
		if show[v - 'a'] == 0 {
			for len(ret) > 0 && count[ret[len(ret) - 1] - 'a'] > 0 && ret[len(ret) - 1] > v {
				show[ret[len(ret) - 1] - 'a'] = 0
				ret = ret[:len(ret) - 1]
			}
			ret = append(ret, v)
			show[v - 'a'] = 1
		}
		count[v - 'a']--
	}

	return string(ret)
}