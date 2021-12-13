package xorOperation

func xorOperation(n int, start int) int {
	i, res := 0, 0
	for i < n {
		res = res ^ (start + 2 * i)
		i++
	}
	return res
}