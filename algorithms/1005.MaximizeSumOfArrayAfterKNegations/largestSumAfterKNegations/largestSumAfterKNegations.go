package largestSumAfterKNegations

import "sort"

func LargestSumAfterKNegations(A []int, K int) int {
	sort.Ints(A)
	res, i := 0, 0

	for K > 0 && i < len(A) {
		if A[i] < 0 {
			A[i] = -A[i]
		} else {
			break
		}
		i++
		K--
	}

	sort.Ints(A)
	if K % 2 != 0 {
		res -= A[0] * 2
	}
	for _, v := range A {
		res += v
	}
	return res
}