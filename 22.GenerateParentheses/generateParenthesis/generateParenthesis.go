package generateParenthesis

func generateParenthesis(n int) []string {
	ret := make([]string, 0)
	var backtrack func(string, int, int, int)
	backtrack = func(cur string, o int, c int, max int) {
		if o == max && c == max {
			ret = append(ret, cur)
			return
		}

		if o < max {
			backtrack(cur + "(", o + 1, c, max)
		}

		if c < o {
			backtrack(cur + ")", o, c + 1, max)
		}
	}
	backtrack("", 0, 0, n)
	return ret
}