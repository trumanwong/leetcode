package fraction

func fraction(cont []int) []int {
	if len(cont) == 1 {
		return []int{cont[0], 1}
	}
	ret := []int{1, cont[len(cont) - 1]}
	for i := len(cont) - 2; i >= 0; i-- {
		numerator, denominator := ret[0], ret[1]
		numerator = numerator + cont[i] * denominator
		ret[0], ret[1] = denominator, numerator
		if i == 0 {
			ret[0], ret[1] = numerator, denominator
		}
	}
	m, n, p := ret[0], ret[1], ret[0] % ret[1]
	for p != 0 {
		m, n = n, p
		p = m % n
	}
	if n != 0 {
		ret[0] /= n
		ret[1] /= n
	}
	return ret
}