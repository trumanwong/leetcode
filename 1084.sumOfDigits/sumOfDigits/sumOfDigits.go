package sumOfDigits

import "sort"

func SumOfDigits(A []int) int {
	sort.Ints(A)
	temp, sum := A[0], 0
	for temp > 0 {
		sum += temp % 10
		temp /= 10
	}
	if sum % 2 == 0 {
		return 1
	}
	return 0
}