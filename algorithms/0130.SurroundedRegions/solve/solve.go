package solve

func Solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			// 从边缘O开始搜索
			isEdge := i == 0 || j == 0 || i == len(board)-1 || j == len(board[i])-1
			if isEdge && board[i][j] == 'O' {
				dfs(board, i, j)
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, i, j int) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) || board[i][j] == 'X' || board[i][j] == '#' {
		return
	}
	board[i][j] = '#'
	dfs(board, i-1, j)
	dfs(board, i+1, j)
	dfs(board, i, j-1)
	dfs(board, i, j+1)
}
