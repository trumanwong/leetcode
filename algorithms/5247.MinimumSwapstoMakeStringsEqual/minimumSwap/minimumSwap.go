package minimumSwap

func MinimumSwap(s1 string, s2 string) int {
	x, y := 0, 0
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			continue
		}

		if s1[i] == 'x' {
			x++
		} else {
				y++
		}
	}

	if (x + y) % 2 == 1 {
		return -1
	}
	res := x / 2 + y / 2
	x %= 2
	y %= 2
	if x != 0 || y != 0 {
		res += 2
	}
	return res
}