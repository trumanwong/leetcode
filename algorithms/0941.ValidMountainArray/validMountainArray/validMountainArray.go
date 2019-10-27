package validMountainArray

func ValidMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}

	i := 0
	for i + 1 < len(A) {
		if A[i + 1] <= A[i] {
			break
		}
		i++
	}
	if i == 0 || i == len(A) - 1 {
		return false
	}
	for i + 1 < len(A) {
		if A[i + 1] >= A[i] {
			return false
		}
		i++
	}
	return true
}