package sortedSquares

func SortedSquares(A []int) []int {
	N := len(A)
	j := 0
	for ; j < N && A[j] < 0; {
		j++
	}
	i := j - 1
	ans := make([]int, N)
	t := 0

	for ; i >= 0 && j < N; {
		i_multi := A[i] * A[i]
		j_multi := A[j] * A[j]
		if i_multi < j_multi {
			ans[t] = i_multi
			i--
		} else {
			ans[t] = j_multi
			j++
		}
		t++
	}
	for ; i >= 0; i-- {
		ans[t] = A[i] * A[i]
		t++
	}

	for ; j < N; j++ {
		ans[t] = A[j] * A[j]
		t++
	}
	return ans
}