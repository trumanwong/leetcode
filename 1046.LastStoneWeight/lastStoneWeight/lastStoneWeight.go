package lastStoneWeight

import "sort"

func LastStoneWeight(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}
	sort.Ints(stones)
	res := 0
	for len(stones) > 1 {
		res = stones[len(stones) - 1] - stones[len(stones) - 2]
		stones = append(stones[0:len(stones) - 2], res)
		sort.Ints(stones)
	}
	return res
}