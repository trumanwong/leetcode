package hammingDistance

func HammingDistance(x int, y int) int {
	x = x ^ y
	res := 0
	for x != 0 {
		res += x & 1
		x >>= 1
	}
	return res
}
