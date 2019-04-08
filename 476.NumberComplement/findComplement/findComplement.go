package findComplement

func FindComplement(num int) int {
	temp, i := num, 0
	for temp > 0 {
		temp >>= 1
		i = (i << 1) + 1
	}
	return num ^ i
}