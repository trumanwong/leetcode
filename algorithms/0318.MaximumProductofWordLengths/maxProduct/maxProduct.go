package maxProduct

func MaxProduct(words []string) int {
	ret := 0
	hash := make([]uint, len(words))
	for i := 0; i < len(words); i++ {
		for _, v := range words[i] {
			hash[i] = hash[i] | (1 << (uint(v) - 'a'))
		}
	}

	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if hash[i]&hash[j] == 0 {
				ret = max(len(words[i])*len(words[j]), ret)
			}
		}
	}
	return ret
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
