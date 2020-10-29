package countOdds

func CountOdds(low int, high int) int {
	return pre(high) - pre(low - 1)
}

func pre(num int) int {
	return (num + 1) >> 1
}