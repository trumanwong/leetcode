package transpose

func Transpose(A [][]int) [][]int {
	col, row := len(A), len(A[0])
	var res [][]int
	for i := 0; i < row; i++ {
		temp := make([]int, col)
		res = append(res, temp)
		for j := 0; j < col; j++ {
			res[i][j] = A[j][i]
		}
	}
	return res
}