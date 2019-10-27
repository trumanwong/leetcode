package arrangeCoins

func ArrangeCoins(n int) int {
	if n <= 1 {
		return n
	}
	start, temp := 0, 1
	for n > 0 {
		start++
		n -= temp
		temp++
	}
	if n < 0{
		return start - 1
	}
	return start
}