package add

func add(a int, b int) int {
	for b != 0 {
		temp := a
		a ^= b // 异或计算未进位的部分
		b = (temp & b) << 1 // 进位部分
	}
	return a
}
