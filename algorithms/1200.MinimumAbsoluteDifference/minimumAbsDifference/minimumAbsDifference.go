package minimumAbsDifference

import "sort"

func MinimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	m := make(map[int][][]int)
	min := 10000000
	for i := 0; i < len(arr) - 1; i++ {
		diff := abs(arr[i] - arr[i + 1])
		if diff > min {
			continue
		}
		m[diff] = append(m[diff], [][]int{{arr[i], arr[i + 1]}}...)
		if diff < min {
			min = diff
		}
	}
	return m[min]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}