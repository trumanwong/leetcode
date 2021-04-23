package setZeroes

func GetZeroes(matrix [][]int) {
	MODIFIED, R, C := -1000000, len(matrix), len(matrix[0])

	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if matrix[r][c] == 0 {
				// We modify the corresponding rows and column elements in place.
				// Note, we only change the non zeroes to MODIFIED

				for k := 0; k < C; k++ {
					if matrix[r][k] != 0 {
						matrix[r][k] = MODIFIED
					}
				}

				for k := 0; k < R; k++ {
					if matrix[k][c] != 0 {
						matrix[k][c] = MODIFIED
					}
				}
			}
		}
	}

	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if matrix[r][c] == MODIFIED {
				matrix[r][c] = 0
			}
		}
	}
}
