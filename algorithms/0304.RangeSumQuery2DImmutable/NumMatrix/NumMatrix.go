package numMatrix

// NumMatrix struct
type NumMatrix struct {
	matrix [][]int
	dp     [][]int
}

// Constructor 构造函数
func Constructor(matrix [][]int) NumMatrix {
	numMatrix := NumMatrix{matrix: matrix}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return numMatrix
	}

	dp := make([][]int, 0)
	for r := 0; r < len(matrix); r++ {
		temp := make([]int, len(matrix[0])+1)
		for c := 0; c < len(matrix[0]); c++ {
			temp[c+1] = temp[c] + matrix[r][c]
		}
		dp = append(dp, temp)
	}
	numMatrix.dp = dp
	return numMatrix
}

// SumRegion 求和
func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	res := 0
	for row := row1; row <= row2; row++ {
		res += this.dp[row][col2+1] - this.dp[row][col1]
	}
	return res
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
