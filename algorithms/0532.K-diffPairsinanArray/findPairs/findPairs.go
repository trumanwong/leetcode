package findPairs

func FindPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}

	diff, m := make(map[int]int), make(map[int]int)
	for _, v := range nums {
		if _, ok := m[v-k]; ok {
			diff[v-k]++
		}
		if _, ok := m[v+k]; ok {
			diff[v]++
		}
		m[v]++
	}
	return len(diff)
}
