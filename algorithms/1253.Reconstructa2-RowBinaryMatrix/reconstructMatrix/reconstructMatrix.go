package reconstructMatrix

func ReconstructMatrix(upper int, lower int, colsum []int) [][]int {
	var ans [][]int
	ans = make([][]int, 2)
	ans[0] = make([]int, len(colsum))
	ans[1] = make([]int, len(colsum))

	for i := 0; i < len(colsum); i++ {
		if colsum[i] == 2 {
			ans[0][i] = 1
			ans[1][i] = 1
			upper--
			lower--
		} else if colsum[i] == 1 {
			if upper >= lower {
				ans[0][i] = 1
				upper--
			} else {
				ans[1][i] = 1
				lower--
			}
		}
	}
	if upper != 0 || lower != 0 {
		return [][]int{}
	}

	return ans
}
