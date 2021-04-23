package rangeBitwiseAnd

func RangeBitwiseAnd(m int, n int) int {
	if m == 0 {
		return 0
	}
	for n > m {
		n = n & (n - 1)
	}
	return n
}
