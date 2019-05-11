package flipAndInvertImage

func FlipAndInvertImage(A [][]int) [][]int {
	for i, arr := range A {
		for j := 0; j < len(arr) / 2; j++ {
			A[i][j], A[i][len(arr) - 1 - j] = A[i][len(arr) - 1 - j], A[i][j]
			if A[i][j] == 0 {
				A[i][j] = 1
			} else {
				A[i][j] = 0
			}
			if A[i][len(arr) - 1 - j] == 0 {
				A[i][len(arr) - 1 - j] = 1
			} else {
				A[i][len(arr) - 1 - j] = 0
			}
		}
		if len(arr) % 2 == 1 {
			if A[i][len(arr) / 2] == 0 {
				A[i][len(arr) / 2] = 1
			} else {
				A[i][len(arr) / 2] = 0
			}
		}
	}
	return A
}