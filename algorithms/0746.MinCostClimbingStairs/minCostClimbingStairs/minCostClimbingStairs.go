package minCostClimbingStairs

func MinCostClimbingStairs(cost []int) int {
	x, y := 0, 0
	for i := 0; i < len(cost); i++ {
		x, y = y, min(x, y) + cost[i]
	}
	return min(x, y)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}