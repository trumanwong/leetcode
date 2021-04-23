package isValidSudoku

import "strconv"

func IsValidSudoku(board [][]byte) bool {
	// 记录某行，某位数字是否已经被摆放
	var row [9][9]bool
	// 记录某列，某位数字是否已经被摆放
	var col [9][9]bool
	// 记录某 3x3 宫格内，某位数字是否已经被摆放
	var block [9][9]bool

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			temp := board[i][j] - 1
			num, _ := strconv.Atoi(string(temp))
			blockIndex := i/3*3 + j/3
			if row[i][num] || col[j][num] || block[blockIndex][num] {
				return false
			} else {
				row[i][num] = true
				col[j][num] = true
				block[blockIndex][num] = true
			}
		}
	}
	return true
}
