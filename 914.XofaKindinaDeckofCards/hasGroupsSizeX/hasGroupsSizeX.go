package hasGroupsSizeX

func hasGroupsSizeX(deck []int) bool {
	m := make(map[int]int)

	X := 10001
	for _, v := range deck {
		m[v]++
	}

	for _, v := range m {
		if v < 2 {
			return false
		}
		if v < X {
			X = v
		}
	}

	for _, v := range m {
		if X != v && gcd(X, v) == 1 {
			return false
		}
	}
	return true
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}

	for b != 0 {
		a, b = b, a % b
	}
	return a
}