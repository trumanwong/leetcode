package maxAreaOfIsland

func MaxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			res = max(dfs(grid, i, j), res)
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func dfs(grid [][]int, r, c int) int {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) || grid[r][c] == 0 {
		return 0
	}
	grid[r][c] = 0
	area := 1
	area += dfs(grid, r, c-1)
	area += dfs(grid, r, c+1)
	area += dfs(grid, r+1, c)
	area += dfs(grid, r-1, c)
	return area
}
