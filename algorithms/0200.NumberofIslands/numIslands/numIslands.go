package numIslands

func NumIslands(grid [][]byte) int {
	res := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == '1' {
				res++
				dfs(grid, r, c)
			}
		}
	}
	return res
}

func dfs(grid [][]byte, r, c int) {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) || grid[r][c] == '0' {
		return
	}
	grid[r][c] = '0'
	dfs(grid, r-1, c)
	dfs(grid, r+1, c)
	dfs(grid, r, c-1)
	dfs(grid, r, c+1)
}
