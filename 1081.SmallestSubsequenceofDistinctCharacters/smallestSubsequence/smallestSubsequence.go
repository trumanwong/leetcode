package smallestSubsequence

func smallestSubsequence(text string) string {
	count, show := make([]int, 26), make([]int, 26)
	for _, ch := range text {
		count[ch - 'a']++
	}

	ret := make([]rune, 0)
	for _, ch := range text {
		if show[ch - 'a'] == 0 {
			for len(ret) > 0 && ch < ret[len(ret) - 1] && count[ret[len(ret) - 1] - 'a'] > 0 {
				show[ret[len(ret) - 1] - 'a'] = 0
				ret = ret[:len(ret) - 1]
			}
			ret = append(ret, ch)
			show[ch - 'a'] = 1
		}
		count[ch - 'a']--
	}
	return string(ret)
}