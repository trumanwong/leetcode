package findLHS

import "math"

func FindLHS(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	res := 0
	for k, v := range m {
		if cur, ok := m[k + 1]; ok {
			res = int(math.Max(float64(cur + v), float64(res)))
		}
	}
	return res
}