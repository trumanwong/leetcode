package minPathSum

func MinPathSum(grid [][]int) int {
	gridRowSize := len(grid)
	gridColSize := len(grid[0])

	var dp [][]int
	for i := 0; i < gridRowSize; i++ {
		tmp := make([]int, gridColSize)
		dp = append(dp, tmp)
	}
	// 初始值为(0, 0)
	dp[0][0] = grid[0][0]
	sum := dp[0][0]
	// 计算左边第一列的所有数字和
	for i := 1; i < gridRowSize; i++ {
		sum += grid[i][0]
		dp[i][0] = sum
	}

	sum = dp[0][0]
	//计算上面第一行的所有数字和
	for i := 1; i < gridColSize; i++ {
		sum += grid[0][i]
		dp[0][i] = sum
	}

	//遍历一遍整个grid，从第二行开始
	for i := 1; i < gridRowSize; i++ {
		//从第2列开始
		for j := 1; j < gridColSize; j++ {
			//对于与任意一点，判断是从左边推进还是从上边推进
			min := dp[i-1][j]
			if dp[i-1][j] > dp[i][j-1] {
				min = dp[i][j-1]
			}
			dp[i][j] = grid[i][j] + min
		}
	}

	//返回推进到的最后一个位置
	return dp[gridRowSize-1][gridColSize-1]
}
