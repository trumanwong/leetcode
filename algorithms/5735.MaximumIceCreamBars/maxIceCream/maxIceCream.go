package maxIceCream

import "sort"

func MaxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	result := 0
	for i := 0; i < len(costs); i++ {
		if coins <= 0 || coins - costs[i] < 0 {
			break
		}
		result++
		coins -= costs[i]
	}
	return result
}
