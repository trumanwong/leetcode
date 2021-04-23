package closedIsland

func ClosedIsland(grid [][]int) int {
	//将边线上的0变成2
	for i := 0; i < len(grid); i++ {
		if grid[i][0] == 0 {
			dfs(grid, i, 0)
		}
		if grid[i][len(grid[0])-1] == 0 {
			dfs(grid, i, len(grid[0])-1)
		}
	}

	for i := 0; i < len(grid[0]); i++ {
		if grid[0][i] == 0 {
			dfs(grid, 0, i)
		}
		if grid[len(grid)-1][i] == 0 {
			dfs(grid, len(grid)-1, i)
		}
	}

	res := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			if grid[i][j] == 0 {
				dfs(grid, i, j)
				res++
			}
		}
	}
	return res
}

func dfs(grid [][]int, x, y int) {
	m, n := len(grid), len(grid[0])
	if x < 0 || y < 0 || x >= m || y >= n || grid[x][y] != 0 {
		return
	}
	grid[x][y] = 2

	dfs(grid, x+1, y)
	dfs(grid, x-1, y)
	dfs(grid, x, y+1)
	dfs(grid, x, y-1)
}
