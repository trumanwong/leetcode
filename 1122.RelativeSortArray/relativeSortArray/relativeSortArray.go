package relativeSortArray

import "sort"

func RelativeSortArray(arr1 []int, arr2 []int) []int {
	ret := []int{}
	sort.Ints(arr1)
	m := make(map[int]int)
	for _, v := range arr1 {
		m[v]++
	}
	for _, v := range arr2 {
		if _, ok := m[v]; ok {
			for i := 0; i < m[v]; i++ {
				ret = append(ret, v)
			}
			m[v] = 0
		}
	}
	for _, v := range arr1 {
		if m[v] == 0 {
			continue
		}

		ret = append(ret, v)
	}
	return ret
}