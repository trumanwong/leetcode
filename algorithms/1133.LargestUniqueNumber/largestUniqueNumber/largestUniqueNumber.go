package largestUniqueNumber

import "sort"

func LargestUniqueNumber(A []int) int {
	ret := -1
	m, a := make(map[int]int), []int{}
	for _, v := range A {
		m[v]++
	}

	for k, v := range m {
		if v == 1 {
			a = append(a, k)
		}
	}

	sort.Ints(a)
	if len(a) > 0 {
		ret = a[len(a)-1]
	}
	return ret
}
