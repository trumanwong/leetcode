package maxScoreSightseeingPair

func MaxScoreSightseeingPair(A []int) int {
	sum, sum_left := 0, 0
	for i, v := range A {
		sum = max(sum, sum_left+v-i)
		sum_left = max(sum_left, v+i)
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
