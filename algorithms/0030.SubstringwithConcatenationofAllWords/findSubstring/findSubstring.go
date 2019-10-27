package findSubstring

func FindSubstring(s string, words []string) []int {
	res := make([]int, 0)
	if len(s) == 0 || len(words) == 0 || len(s) < len(words[0]) * len(words) {
		return res
	}

	return res
}