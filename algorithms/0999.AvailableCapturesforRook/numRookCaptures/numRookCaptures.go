package numRookCaptures

func NumRookCaptures(board [][]byte) int {
	x, y := 0, 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'R' {
				x, y = i, j
				break
			}
		}
	}

	res := 0
	for i := x; i < len(board[0]); i++ {
		if board[i][y] == 'B' {
			break
		}
		if board[i][y] == 'p' {
			res++
			break
		}
	}
	for i := x; i >= 0; i-- {
		if board[i][y] == 'B' {
			break
		}
		if board[i][y] == 'p' {
			res++
			break
		}
	}
	for i := y; i < len(board); i++ {
		if board[x][i] == 'B' {
			break
		}
		if board[x][i] == 'p' {
			res++
			break
		}
	}
	for i := y; i >= 0; i-- {
		if board[x][i] == 'B' {
			break
		}
		if board[x][i] == 'p' {
			res++
			break
		}
	}
	return res
}
