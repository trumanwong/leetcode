package exist

func Exist(board [][]byte, word string) bool {
	// 深度优先搜索
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, i, j, start int) bool {
	if start >= len(word) {
		return true
	}

	if  i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[start] {
		return false
	}

	board[i][j] = '#' // 标记被访问过的字母
	if dfs(board, word, i + 1, j, start + 1) || dfs(board, word, i - 1, j, start + 1) || dfs(board, word, i, j + 1, start + 1) || dfs(board, word, i, j - 1, start + 1) {
		return true
	}
	board[i][j] = word[start]
	return false
}