package findWords

func FindWords(board [][]byte, words []string) []string {
	res := make([]string, 0)
	var dfs func(word string, row, col, start int) bool
	dfs = func(word string, row, col, start int) bool {
		if start == len(word) {
			return true
		}

		if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) || board[row][col] != word[start] {
			return false
		}

		board[row][col] = '#'
		start++
		if dfs(word, row+1, col, start) || dfs(word, row-1, col, start) || dfs(word, row, col+1, start) || dfs(word, row, col-1, start) {
			board[row][col] = word[start-1]
			return true
		}
		board[row][col] = word[start-1]
		return false
	}
	for i := 0; i < len(words); i++ {
		temp := false
		for row := 0; row < len(board); row++ {
			for col := 0; col < len(board[0]); col++ {
				if dfs(words[i], row, col, 0) {
					temp = true
					break
				}
			}
			if temp {
				res = append(res, words[i])
				break
			}
		}
	}
	return res
}
