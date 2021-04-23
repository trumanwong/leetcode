package smallestRangeI

func SmallestRangeI(A []int, K int) int {
	length := len(A)
	if length == 1 {
		return 0
	}

	min := A[0]
	max := A[0]
	for i := 1; i < length; i++ {
		if min > A[i] {
			min = A[i]
		}
		if max < A[i] {
			max = A[i]
		}
	}

	diff := max - min
	if diff > K {
		if diff-2*K < 0 {
			return 0
		}
		return diff - 2*K
	}
	return 0
}
