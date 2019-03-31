package rotate

func Rotate(matrix [][]int)  {
	lenMatrix := len(matrix)
	// 上下翻转
	for i := 0; i < lenMatrix/2; i++ {
		matrix[i], matrix[lenMatrix - i - 1] = matrix[lenMatrix - i - 1], matrix[i]
	}

	// 对角线翻转
	for j := 0; j < lenMatrix; j++ {
		for k := j; k < lenMatrix; k++ {
			matrix[j][k], matrix[k][j] = matrix[k][j], matrix[j][k]
		}
	}
}