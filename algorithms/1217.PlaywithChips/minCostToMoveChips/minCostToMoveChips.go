package minCostToMoveChips

func MinCostToMoveChips(chips []int) int {
	n, cnt := len(chips), 0
	for _, v := range chips {
		cnt = cnt + v%2
	}
	return min(cnt, n-cnt)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
