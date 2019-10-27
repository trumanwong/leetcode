package balancedStringSplit

func BalancedStringSplit(s string) int {
	res := 0
	r, l := 0, 0
	for _, v := range s {
		if v == 'R' {
			r++
		} else {
			l++
		}
		if r == l {
			r, l = 0, 0
			res++
		}
	}
	if r > 0 && r == l {
		res++
	}
	return res
}