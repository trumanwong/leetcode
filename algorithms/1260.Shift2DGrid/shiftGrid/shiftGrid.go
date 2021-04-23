package shiftGrid

func ShiftGrid(grid [][]int, k int) [][]int {
	for k > 0 {
		temp := make([][]int, 0)
		for i := 0; i < len(grid); i++ {
			t := make([]int, len(grid[i]))
			copy(t, grid[i])
			temp = append(temp, t)
		}
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i])-1; j++ {
				grid[i][j+1] = temp[i][j]
			}
			if i < len(temp)-1 {
				grid[i+1][0] = temp[i][len(grid[0])-1]
			}
		}
		grid[0][0] = temp[len(grid)-1][len(grid[0])-1]
		k--
	}
	return grid
}
