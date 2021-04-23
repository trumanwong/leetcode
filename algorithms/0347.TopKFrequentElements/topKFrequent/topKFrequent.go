package topKFrequent

import "sort"

func TopKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	ret, counts := make([]int, 0, k), make([]int, 0, len(m))
	for _, count := range m {
		counts = append(counts, count)
	}
	sort.Ints(counts)
	limit := counts[len(counts)-k]
	for k, v := range m {
		if v >= limit {
			ret = append(ret, k)
		}
	}
	return ret
}
