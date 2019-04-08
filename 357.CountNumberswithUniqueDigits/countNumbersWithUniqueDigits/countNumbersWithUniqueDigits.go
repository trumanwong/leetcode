package countNumbersWithUniqueDigits

func CountNumbersWithUniqueDigits(n int) int {
	res := 0
	for i := 0; i < n; i++ {
		temp := 9
		for j := 0; j < i; j++ {
			temp *= 9 - j
		}
		res += temp
	}
	return res + 1
}