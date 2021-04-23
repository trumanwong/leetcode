package spiralOrder

func SpiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	if len(matrix) == 0 {
		return res
	}

	rowRight, rowLeft, colRight, colLeft := 0, len(matrix)-1, 0, len(matrix[0])-1
	for rowRight <= rowLeft && colRight <= colLeft {
		for i := colRight; i <= colLeft; i++ {
			res = append(res, matrix[rowRight][i])
		}

		for i := rowRight + 1; i <= rowLeft; i++ {
			res = append(res, matrix[i][colLeft])
		}

		if rowRight < rowLeft && colRight < colLeft {
			for i := colLeft - 1; i > colRight; i-- {
				res = append(res, matrix[rowLeft][i])
			}

			for i := rowLeft; i > rowRight; i-- {
				res = append(res, matrix[i][colRight])
			}
		}
		rowRight++
		rowLeft--
		colRight++
		colLeft--
	}
	return res
}
