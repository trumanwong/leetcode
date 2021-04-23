package numPairsDivisibleBy60

func NumPairsDivisibleBy60(time []int) int {
	needs := make([]int, 60)
	res := 0
	for _, v := range time {
		if v%60 == 0 {
			res += needs[0]
		} else {
			res += needs[60-v%60]
		}
		needs[v%60] += 1
	}
	return res
}
