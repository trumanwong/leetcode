package subtractProductAndSum

func SubtractProductAndSum(n int) int {
	sum, mul := 0, 1
	for n != 0 {
		sum += n % 10
		mul *= n % 10
		n /= 10
	}
	return mul - sum
}
