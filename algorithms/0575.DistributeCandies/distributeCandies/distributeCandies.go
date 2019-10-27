package distributeCandies

func DistributeCandies(candies []int) int {
	m := make(map[int]int)
	for _, v := range candies {
		m[v]++
	}

	if len(candies) / 2 < len(m) {
		return len(candies) / 2
	}
	return len(m)
}