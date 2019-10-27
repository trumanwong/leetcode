package letterCombinations

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	bDigits := []byte(digits)
	var recursive func(index int, b []byte)
	m, res := make(map[string]int), []string{}
	nums := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	recursive = func(index int, b []byte) {
		if len(bDigits) == len(b) {
			m[string(b)]++
			return
		}

		temp := nums[bDigits[index] - '0' - 2]
		for _, v := range temp {
			b = append(b, byte(v))
			recursive(index + 1, b)
			b = b[0:len(b) - 1]
		}
	}
	recursive(0, []byte{})
	for k, _ := range m {
		res = append(res, k)
	}
	return res
}