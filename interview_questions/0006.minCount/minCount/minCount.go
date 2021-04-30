package minCount

func minCount(coins []int) int {
	res := 0
	for _, v := range coins {
		res += v / 2 + v % 2
	}
	return res
}