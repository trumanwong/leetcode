package getSum

func GetSum(a int, b int) int {
	if b == 0 {
		return a
	}
	return GetSum(a ^ b, (a & b) << 1)
}