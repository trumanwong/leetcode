package isPowerOfFour

func IsPowerOfFour(num int) bool {
	if num < 0 || (num & (num-1)) != 0 {
		return false
	}
	// 4的幂次-1一定是3的倍数
	return (num - 1) % 3 == 0
}