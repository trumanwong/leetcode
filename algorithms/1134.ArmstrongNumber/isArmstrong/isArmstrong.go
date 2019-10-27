package isArmstrong

import "strconv"

func IsArmstrong(N int) bool {
	sum, temp := 0, N
	str := strconv.Itoa(N)
	for temp > 0 {
		digit := temp % 10
		sum += pow(digit, len(str))
		temp = temp / 10
	}
	if sum == N {
		return true
	}
	return false
}

func pow(n, N int) int {
	ret := 1
	for i := 0; i < N; i++ {
		ret = ret * n
	}
	return ret
}