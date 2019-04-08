package maxWidthRamp

func MaxWidthRamp(A []int) int {
	res := 0
	for i := 0; i < len(A); i++ {
		for j := len(A) - 1; j > i; j-- {
			if j - i < res {
				break
			}
			if A[i] <= A[j] && j - i > res {
				res = j - i
			}
		}
	}
	return res
}