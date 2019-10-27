package findAnagrams

func FindAnagrams(s string, p string) []int {
	sLen, pLen, res := len(s), len(p), []int{}
	mul, end := 1, sLen - pLen
	for _, v := range p {
		mul *= int(v)
	}

	for i := 0; i <= end; i++ {
		if getMul(s[i:i + pLen]) == mul {
			res = append(res, i)
		}
	}
	return res
}

func getMul(s string) int {
	mul := 1
	for _, v := range s {
		mul *= int(v)
	}
	return mul
}