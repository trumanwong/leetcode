package lastRemaining

func LastRemaining(n int) int {
	if n == 1 {
		return 1
	} else {
		return 2 * (n/2 + 1 - LastRemaining(n/2))
	}
}
