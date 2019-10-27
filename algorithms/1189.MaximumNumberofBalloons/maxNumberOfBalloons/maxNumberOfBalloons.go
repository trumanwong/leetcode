package maxNumberOfBalloons

func maxNumberOfBalloons(text string) int {
	// balloon
	b, a, l, o, n := 0, 0, 0, 0, 0
	ret := 10000
	for _, v := range text {
		if v == 'b' {
			b++
		} else if v == 'a' {
			a++
		} else if v == 'l' {
			l++
		} else if v == 'o' {
			o++
		} else if v == 'n' {
			n++
		}
	}
	l, o = l / 2, o /2
	ret = min(ret, min(b, min(a, min(l, min(o, n)))))
	return ret
}

func min (x, y int) int {
	if x < y {
		return x
	}
	return y
}