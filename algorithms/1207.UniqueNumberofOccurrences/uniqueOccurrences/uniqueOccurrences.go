package uniqueOccurrences

func UniqueOccurrences(arr []int) bool {
	counts, nums := make(map[int]int), make(map[int]int)
	for _, v := range arr {
		nums[v]++
	}

	for k, v := range nums {
		if _, ok := counts[v]; ok {
			return false
		}
		counts[v] = k
	}
	return true
}