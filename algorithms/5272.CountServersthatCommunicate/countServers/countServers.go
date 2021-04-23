package countServers

func CountServers(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			u, d, l, r := i-1, i+1, j-1, j+1
			judge := false
			for u >= 0 && !judge {
				if grid[u][j] == 1 {
					judge = true
					res++
				}
				u--
			}
			for d < len(grid) && !judge {
				if grid[d][j] == 1 {
					judge = true
					res++
				}
				d++
			}
			for l >= 0 && !judge {
				if grid[i][l] == 1 {
					judge = true
					res++
				}
				l--
			}
			for r < len(grid[i]) && !judge {
				if grid[i][r] == 1 {
					judge = true
					res++
				}
				r++
			}
		}
	}
	return res
}
