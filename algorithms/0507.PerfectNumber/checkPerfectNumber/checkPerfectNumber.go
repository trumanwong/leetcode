package checkPerfectNumber

func CheckPerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}

	sum := 1
	for i := 2; i * i <= num; i++ {
		if num % i == 0 {
			sum = sum + i + num / i
		}
	}

	return sum == num
}