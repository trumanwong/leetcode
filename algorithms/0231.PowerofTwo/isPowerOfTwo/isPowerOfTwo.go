package isPowerOfTwo

func IsPowerOfTwo(n int) bool {
	// &运算，同1则1。
	// 2的幂次方在二进制下，只有1位是1，其余全是0。
	return n > 0 && (n&-n == n)
}
