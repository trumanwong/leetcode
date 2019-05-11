package rotatedDigits

func RotatedDigits(N int) int {
	res := 0
	for i := 1; i <= N; i++ {
		temp := i
		isNotNumber, isChange := false, false
		for temp != 0 {
			mod := temp % 10
			temp /= 10
			if mod == 2 || mod == 5 || mod == 6 || mod == 9 {
				isChange = true
			} else if mod != 0 && mod != 1 && mod != 8 {
				isNotNumber = true
				break
			}
		}
		if !isNotNumber && isChange {
			res++
		}
	}
	return res
}