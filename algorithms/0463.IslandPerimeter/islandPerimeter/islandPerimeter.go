package islandPerimeter

func IslandPerimeter(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				res += 4
				if j > 0 && grid[i][j - 1] == 1 {
					res -= 2
				}
				if i > 0 && grid[i - 1][j] == 1 {
					res -= 2
				}
			}
		}
	}
	return res
}