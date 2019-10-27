package arraysIntersection

import "sort"

func ArraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	m := make(map[int]int)
	for _, v := range arr1 {
		m[v]++
	}

	for _, v := range arr2 {
		m[v]++
	}

	for _, v := range arr3 {
		m[v]++
	}

	res := make([]int, 0)
	for k, v := range m {
		if v == 3 {
			res = append(res, k)
		}
	}
	sort.Ints(res)
	return res
}