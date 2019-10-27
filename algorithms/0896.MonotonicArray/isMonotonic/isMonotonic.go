package isMonotonic

func IsMonotonic(A []int) bool {
	isIncreasing, isDecreasing := true, true
	i := 0
	for i < len(A) - 1 {
		if A[i] > A[i + 1] {
			isIncreasing = false
		} else if A[i] < A[i + 1] {
			isDecreasing = false
		}
		i++
	}
	return isIncreasing || isDecreasing
}