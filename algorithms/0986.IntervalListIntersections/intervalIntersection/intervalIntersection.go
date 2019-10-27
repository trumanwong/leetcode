package intervalIntersection

func IntervalIntersection(A [][]int, B [][]int) [][]int {
	res := make([][]int, 0)
	i, j := 0, 0
	for i < len(A) && j < len(B) {
		l, r := max(A[i][0], B[j][0]), min(A[i][1], B[j][1])
		if l <= r {
			res = append(res, []int{l, r})
		}

		if A[i][1] < B[j][1] {
			i++
		} else {
			j++
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}