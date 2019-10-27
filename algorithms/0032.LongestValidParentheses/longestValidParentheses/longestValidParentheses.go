package longestValidParentheses

func LongestValidParentheses(s string) int {
	l, r, maxLength := 0, 0, 0
	for _, v := range s {
		if v == '(' {
			l++
		} else {
			r++
		}

		if l == r {
			maxLength = max(maxLength, 2*r)
		} else if r >= l {
			l, r = 0, 0
		}
	}

	l, r = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			l++
		} else {
				r++
		}

		if l == r {
			maxLength = max(maxLength, 2 * l)
		} else if l >= r {
			l, r = 0, 0
		}
	}
	return maxLength
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}