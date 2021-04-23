package twoSumLessThanK

func TwoSumLessThanK(A []int, K int) int {
	ret := -1
	for i := 0; i < len(A); i++ {
		if A[i] >= K {
			continue
		}
		for j := i + 1; j < len(A); j++ {
			if A[j] >= K {
				continue
			}
			sum := A[i] + A[j]
			if ret < sum && sum < K {
				ret = sum
			}
		}
	}
	return ret
}
