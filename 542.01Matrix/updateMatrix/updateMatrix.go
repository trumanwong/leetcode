package updateMatrix

func UpdateMatrix(matrix [][]int) [][]int {
	row, col := len(matrix), len(matrix[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 1 {
				matrix[i][j] = row + col
			}

			if i > 0 {
				matrix[i][j] = min(matrix[i][j], matrix[i - 1][j] + 1)
			}
			if j > 0 {
				matrix[i][j] = min(matrix[i][j], matrix[i][j - 1] + 1)
			}
		}
	}

	for i := row - 1; i >= 0; i-- {
		for j := col - 1; j >= 0; j-- {
			if i < row - 1 {
				matrix[i][j] = min(matrix[i][j], matrix[i + 1][j] + 1)
			}
			if j < col - 1 {
				matrix[i][j] = min(matrix[i][j], matrix[i][j + 1] + 1)
			}
		}
	}

	return matrix
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}