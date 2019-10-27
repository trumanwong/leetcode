package rearrangeBarcodes

import "sort"

func RearrangeBarcodes(barcodes []int) []int {
	m, keys := make(map[int]int), []int{}
	for _, v := range barcodes {
		m[v]++
		if m[v] == 1 {
			keys = append(keys, v)
		}
	}

	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	res := make([]int, len(barcodes))
	start, index := 0, 0
	for index < len(keys) {
		k := keys[index]
		for i := start; i < len(res); i += 2 {
			res[i] = k
			m[k]--
			if m[k] == 0 {
				index++
				if index >= len(keys) {
					break
				}
				k = keys[index]
			}
		}
		start++
	}
	return res
}