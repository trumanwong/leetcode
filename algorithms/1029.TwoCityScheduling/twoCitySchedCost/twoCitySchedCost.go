package twoCitySchedCost

import "sort"

func TwoCitySchedCost(costs [][]int) int {
	diff, res := make([]int, len(costs)), 0
	for i, v := range costs {
		diff[i] = v[1] - v[0]
		res += v[0]
	}

	sort.Ints(diff)
	for i := 0; i < len(costs) / 2; i++ {
		res += diff[i]
	}
	return res
}