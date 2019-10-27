package hammingWeight

func HammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		res += int(num & 1)
		num = num >> 1
	}
	return res
}